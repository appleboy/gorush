package notify

import (
	"context"
	"crypto/ecdsa"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	apns "github.com/sideshow/apns2"

	"github.com/eencloud/gorush/config"
	"github.com/eencloud/gorush/core"
	"github.com/eencloud/gorush/logx"
	"github.com/eencloud/gorush/status"

	"github.com/eencloud/goeen/dhash"

	"github.com/mitchellh/mapstructure"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"
	"golang.org/x/net/http2"
)

var (
	idleConnTimeout = 90 * time.Second
	tlsDialTimeout  = 20 * time.Second
	tcpKeepAlive    = 60 * time.Second
)

const (
	dotP8  = ".p8"
	dotPEM = ".pem"
	dotP12 = ".p12"
)

var doOnce sync.Once

// DialTLS is the default dial function for creating TLS connections for
// non-proxied HTTPS requests.
var DialTLS = func(cfg *tls.Config) func(network, addr string) (net.Conn, error) {
	return func(network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{
			Timeout:   tlsDialTimeout,
			KeepAlive: tcpKeepAlive,
		}
		return tls.DialWithDialer(dialer, network, addr, cfg)
	}
}

// Sound sets the aps sound on the payload.
type Sound struct {
	Critical int     `json:"critical,omitempty"`
	Name     string  `json:"name,omitempty"`
	Volume   float32 `json:"volume,omitempty"`
}

type Watch struct {
	*dhash.WatchValue
	values []string
}

func RemoveToken(token string, esn string, device_type string) {

	for i := 0; i < 5; i++ { // Try up to 5 times
		ret := NewWatch(esn, device_type)
		if ret == nil {
			logx.LogError.Error("Could not get watch for %s", esn)
		}

		err := ret.Get()

		if err == dhash.WatchEmpty {
			logx.LogError.Error("Attempted to remove token from null list")
		} else if err != nil {
			logx.LogError.Error("Dhash Error: ", err.Error())
			continue
		} else {
			b := ret.values[:0]
			for i, x := range ret.values {
				logx.LogError.Error("Scanning value %s", x)
				if x == token {
					ret.values = append(ret.values[:i], ret.values[i+1:]...)
					break
				}
			}

			ret.values = b

			saveErr := ret.Save()

			if saveErr != nil { // data most likely changed, retry the process
				logx.LogError.Error("Save Error: ", saveErr.Error())
				continue
			} else {
				log.Printf("Saved string array: %v", ret.values)
				break
			}
		}
	}
}

func NewWatch(esn string, device_type string) *Watch {
	w := &Watch{}
	str := fmt.Sprintf("com.eencloud.push_tokens.%s.%s", esn, device_type)
	dh, err := dhash.Resolve(str)
	log.Printf("dhash key: %s", str)
	if err != nil {
		logx.LogError.Error("Failed to resolve dhash token", err.Error())
		return nil
	}

	w.WatchValue = dhash.NewWatchValue(dh, &w.values, str)
	return w
}

// InitAPNSClient use for initialize APNs Client.
func InitAPNSClient(ctx context.Context, cfg *config.ConfYaml) error {
	if cfg.Ios.Enabled {
		var err error
		var authKey *ecdsa.PrivateKey
		var certificateKey tls.Certificate
		var ext string

		var addr = os.Getenv("EEN_DHASH_ADDRESS")
		log.Printf("Initializing dhash service at %s", addr)

		dhash.Initialize(addr)

		if cfg.Ios.KeyPath != "" {
			ext = filepath.Ext(cfg.Ios.KeyPath)

			switch ext {
			case dotP12:
				certificateKey, err = certificate.FromP12File(cfg.Ios.KeyPath, cfg.Ios.Password)
			case dotPEM:
				certificateKey, err = certificate.FromPemFile(cfg.Ios.KeyPath, cfg.Ios.Password)
			case dotP8:
				authKey, err = token.AuthKeyFromFile(cfg.Ios.KeyPath)
			default:
				err = errors.New("wrong certificate key extension")
			}

			if err != nil {
				logx.LogError.Error("Cert Error:", err.Error())

				return err
			}
		} else if cfg.Ios.KeyBase64 != "" {
			ext = "." + cfg.Ios.KeyType
			key, err := base64.StdEncoding.DecodeString(cfg.Ios.KeyBase64)
			if err != nil {
				logx.LogError.Error("base64 decode error:", err.Error())

				return err
			}
			switch ext {
			case dotP12:
				certificateKey, err = certificate.FromP12Bytes(key, cfg.Ios.Password)
			case dotPEM:
				certificateKey, err = certificate.FromPemBytes(key, cfg.Ios.Password)
			case dotP8:
				authKey, err = token.AuthKeyFromBytes(key)
			default:
				err = errors.New("wrong certificate key type")
			}

			if err != nil {
				logx.LogError.Error("Cert Error:", err.Error())

				return err
			}
		}

		if ext == dotP8 {
			if cfg.Ios.KeyID == "" || cfg.Ios.TeamID == "" {
				msg := "you should provide ios.KeyID and ios.TeamID for p8 token"
				logx.LogError.Error(msg)
				return errors.New(msg)
			}
			token := &token.Token{
				AuthKey: authKey,
				// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
				KeyID: cfg.Ios.KeyID,
				// TeamID from developer account (View Account -> Membership)
				TeamID: cfg.Ios.TeamID,
			}

			ApnsClient, err = newApnsTokenClient(cfg, token)
		} else {
			ApnsClient, err = newApnsClient(cfg, certificateKey)
		}

		if h2Transport, ok := ApnsClient.HTTPClient.Transport.(*http2.Transport); ok {
			configureHTTP2ConnHealthCheck(h2Transport)
		}

		if err != nil {
			logx.LogError.Error("Transport Error:", err.Error())

			return err
		}

		doOnce.Do(func() {
			MaxConcurrentIOSPushes = make(chan struct{}, cfg.Ios.MaxConcurrentPushes)
		})
	}

	return nil
}

func newApnsClient(cfg *config.ConfYaml, certificate tls.Certificate) (*apns2.Client, error) {
	var client *apns2.Client

	if cfg.Ios.Production {
		client = apns2.NewClient(certificate).Production()
	} else {
		client = apns2.NewClient(certificate).Development()
	}

	if cfg.Core.HTTPProxy == "" {
		return client, nil
	}

	//nolint:gosec
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}

	if len(certificate.Certificate) > 0 {
		//nolint:staticcheck
		tlsConfig.BuildNameToCertificate()
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		DialTLS:         DialTLS(tlsConfig),
		Proxy:           http.DefaultTransport.(*http.Transport).Proxy,
		IdleConnTimeout: idleConnTimeout,
	}

	h2Transport, err := http2.ConfigureTransports(transport)
	if err != nil {
		return nil, err
	}

	configureHTTP2ConnHealthCheck(h2Transport)

	client.HTTPClient.Transport = transport

	return client, nil
}

func newApnsTokenClient(cfg *config.ConfYaml, token *token.Token) (*apns2.Client, error) {
	var client *apns2.Client

	if cfg.Ios.Production {
		client = apns2.NewTokenClient(token).Production()
	} else {
		client = apns2.NewTokenClient(token).Development()
	}

	if cfg.Core.HTTPProxy == "" {
		return client, nil
	}

	transport := &http.Transport{
		DialTLS:         DialTLS(nil),
		Proxy:           http.DefaultTransport.(*http.Transport).Proxy,
		IdleConnTimeout: idleConnTimeout,
	}

	h2Transport, err := http2.ConfigureTransports(transport)
	if err != nil {
		return nil, err
	}

	configureHTTP2ConnHealthCheck(h2Transport)

	client.HTTPClient.Transport = transport

	return client, nil
}

func configureHTTP2ConnHealthCheck(h2Transport *http2.Transport) {
	h2Transport.ReadIdleTimeout = 1 * time.Second
	h2Transport.PingTimeout = 1 * time.Second
}

func iosAlertDictionary(notificationPayload *payload.Payload, req *PushNotification) *payload.Payload {
	// Alert dictionary

	if len(req.Title) > 0 {
		notificationPayload.AlertTitle(req.Title)
	}

	if len(req.InterruptionLevel) > 0 {
		notificationPayload.InterruptionLevel(payload.EInterruptionLevel(req.InterruptionLevel))
	}

	if len(req.Message) > 0 && len(req.Title) > 0 {
		notificationPayload.AlertBody(req.Message)
	}

	if len(req.Alert.Title) > 0 {
		notificationPayload.AlertTitle(req.Alert.Title)
	}

	// Apple Watch & Safari display this string as part of the notification interface.
	if len(req.Alert.Subtitle) > 0 {
		notificationPayload.AlertSubtitle(req.Alert.Subtitle)
	}

	if len(req.Alert.TitleLocKey) > 0 {
		notificationPayload.AlertTitleLocKey(req.Alert.TitleLocKey)
	}

	if len(req.Alert.LocArgs) > 0 {
		notificationPayload.AlertLocArgs(req.Alert.LocArgs)
	}

	if len(req.Alert.TitleLocArgs) > 0 {
		notificationPayload.AlertTitleLocArgs(req.Alert.TitleLocArgs)
	}

	if len(req.Alert.Body) > 0 {
		notificationPayload.AlertBody(req.Alert.Body)
	}

	if len(req.Alert.LaunchImage) > 0 {
		notificationPayload.AlertLaunchImage(req.Alert.LaunchImage)
	}

	if len(req.Alert.LocKey) > 0 {
		notificationPayload.AlertLocKey(req.Alert.LocKey)
	}

	if len(req.Alert.Action) > 0 {
		notificationPayload.AlertAction(req.Alert.Action)
	}

	if len(req.Alert.ActionLocKey) > 0 {
		notificationPayload.AlertActionLocKey(req.Alert.ActionLocKey)
	}

	// General
	if len(req.Category) > 0 {
		notificationPayload.Category(req.Category)
	}

	if len(req.Alert.SummaryArg) > 0 {
		notificationPayload.AlertSummaryArg(req.Alert.SummaryArg)
	}

	if req.Alert.SummaryArgCount > 0 {
		notificationPayload.AlertSummaryArgCount(req.Alert.SummaryArgCount)
	}

	if len(req.ContentState) > 0 {
		notificationPayload.SetContentState(req.ContentState)
	}

	if req.StaleDate > 0 {
		notificationPayload.SetStaleDate(req.StaleDate)
	}

	if req.DismissalDate > 0 {
		notificationPayload.SetDismissalDate(req.DismissalDate)
	}

	if len(req.Event) > 0 {
		notificationPayload.SetEvent(payload.ELiveActivityEvent(req.Event))
	}

	if req.Timestamp > 0 {
		notificationPayload.SetTimestamp(req.Timestamp)
	}

	return notificationPayload
}

// GetIOSNotification use for define iOS notification.
// The iOS Notification Payload (Payload Key Reference)
// Ref: https://apple.co/2VtH6Iu
func GetIOSNotification(req *PushNotification) *apns2.Notification {
	notification := &apns2.Notification{
		ApnsID:     req.ApnsID,
		Topic:      req.Topic,
		CollapseID: req.CollapseID,
	}

	if req.Expiration != nil {
		notification.Expiration = time.Unix(*req.Expiration, 0)
	}

	if len(req.Priority) > 0 {
		if req.Priority == "normal" {
			notification.Priority = apns2.PriorityLow
		} else if req.Priority == HIGH {
			notification.Priority = apns2.PriorityHigh
		}
	}

	if len(req.PushType) > 0 {
		notification.PushType = apns2.EPushType(req.PushType)
	}

	payload := payload.NewPayload()

	// add alert object if message length > 0 and title is empty
	if len(req.Message) > 0 && req.Title == "" {
		payload.Alert(req.Message)
	}

	// zero value for clear the badge on the app icon.
	if req.Badge != nil && *req.Badge >= 0 {
		payload.Badge(*req.Badge)
	}

	if req.MutableContent == 1 {
		payload.MutableContent()
	}

	switch req.Sound.(type) {
	// from http request binding
	case map[string]interface{}:
		result := &Sound{}
		_ = mapstructure.Decode(req.Sound, &result)
		payload.Sound(result)
	// from http request binding for non critical alerts
	case string:
		payload.Sound(&req.Sound)
	case Sound:
		payload.Sound(&req.Sound)
	}

	if len(req.SoundName) > 0 {
		payload.SoundName(req.SoundName)
	}

	if req.SoundVolume > 0 {
		payload.SoundVolume(req.SoundVolume)
	}

	if req.ContentAvailable {
		payload.ContentAvailable()
	}

	if len(req.URLArgs) > 0 {
		payload.URLArgs(req.URLArgs)
	}

	if len(req.ThreadID) > 0 {
		payload.ThreadID(req.ThreadID)
	}

	for k, v := range req.Data {
		payload.Custom(k, v)
	}

	payload = iosAlertDictionary(payload, req)

	notification.Payload = payload

	return notification
}

func getApnsClient(cfg *config.ConfYaml, req *PushNotification) (client *apns2.Client) {
	switch {
	case req.Production:
		client = ApnsClient.Production()
	case req.Development:
		client = ApnsClient.Development()
	default:
		if cfg.Ios.Production {
			client = ApnsClient.Production()
		} else {
			client = ApnsClient.Development()
		}
	}

	return
}

// PushToIOS provide send notification to APNs server.
func PushToIOS(ctx context.Context, req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for iOS")

	var (
		retryCount = 0
		maxRetry   = cfg.Ios.MaxRetry
	)

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	resp = &ResponsePush{}

Retry:
	var newTokens []string

	notification := GetIOSNotification(req)
	client := getApnsClient(cfg, req)

	// Log batch information
	tokenCount := len(req.Tokens)
	logx.LogAccess.Infof("Processing iOS push batch: %d tokens (retry attempt: %d/%d)", tokenCount, retryCount, maxRetry)

	// Track batch results
	var successCount, errorCount, nilResponseCount int32

	var wg sync.WaitGroup
	for i := 0; i < tokenCount; i++ {
		token := req.Tokens[i]
		userId := req.UserIds[i]
		// occupy push slot
		MaxConcurrentIOSPushes <- struct{}{}
		wg.Add(1)
		go func(notification apns2.Notification, token string) {
			notification.DeviceToken = token
			// send ios notification
			res, err := client.PushWithContext(ctx, &notification)
			if err != nil || (res != nil && res.StatusCode != http.StatusOK) {
				// Handle nil response case (network errors, timeouts, context cancellation)
				if res == nil {
					atomic.AddInt32(&nilResponseCount, 1)
					logx.LogError.Errorf("iOS push failed with nil response for token %s, userId %s: %v", token, userId, err)
					// Log when we skip token removal due to nil response
					logx.LogError.Warnf("Skipping token removal check for token %s (nil response, likely network/timeout error)", token)
				} else {
					// APNS returned an error response
					atomic.AddInt32(&errorCount, 1)

					// Use APNS reason as error message
					if err == nil {
						// error message:
						// ref: https://github.com/sideshow/apns2/blob/master/response.go#L14-L65
						err = errors.New(res.Reason)
					}

					// Check if we should retry (server errors only)
					if res.StatusCode >= http.StatusInternalServerError {
						newTokens = append(newTokens, token)
					}

					// Check if token should be removed (bad token errors)
					reasons := []string{apns.ReasonBadDeviceToken, apns.ReasonDeviceTokenNotForTopic, apns.ReasonUnregistered}
					for _, a := range reasons {
						if a == res.Reason {
							go RemoveToken(token, userId, "ios")
							break
						}
					}
				}

				// Log push failure
				errLog := logPush(cfg, core.FailedPush, token, userId, req, err)
				if notifID, _ := req.Data["notification_id"].(string); notifID != "" {
					logx.LogAccess.Infof("iOS push failed | notification_id: %s | error: %v", notifID, err)
				}
				resp.Logs = append(resp.Logs, errLog)
				status.StatStorage.AddIosError(1)
			}

			if res != nil && res.Sent() {
				atomic.AddInt32(&successCount, 1)
				logPush(cfg, core.SucceededPush, token, userId, req, nil)
				notifID, _ := req.Data["notification_id"].(string)
				if notifID != "" {
					logx.LogAccess.Infof("iOS push sent | notification_id: %s | apns_id_req: %s | apns_id_res: %s | sent_at: %s",
						notifID, req.ApnsID, res.ApnsID, time.Now().UTC().Format(time.RFC3339))
				} else {
					logx.LogAccess.Infof("iOS push sent | apns_id_req: %s | apns_id_res: %s | sent_at: %s",
						req.ApnsID, res.ApnsID, time.Now().UTC().Format(time.RFC3339))
				}
				status.StatStorage.AddIosSuccess(1)
			}

			// free push slot
			<-MaxConcurrentIOSPushes
			wg.Done()
		}(*notification, token)
	}

	wg.Wait()

	// Log batch summary (use atomic loads for proper memory ordering)
	logx.LogAccess.Infof("iOS push batch complete: %d tokens processed | Success: %d | APNS errors: %d | Network/timeout errors: %d | Retrying: %d",
		tokenCount, atomic.LoadInt32(&successCount), atomic.LoadInt32(&errorCount), atomic.LoadInt32(&nilResponseCount), len(newTokens))

	if len(newTokens) > 0 && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return resp, nil
}
