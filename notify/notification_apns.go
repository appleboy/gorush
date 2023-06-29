package notify

import (
	"crypto/ecdsa"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"net"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"

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

// InitAPNSClient use for initialize APNs Client.
func InitAPNSClient(tenantId string, tenant config.SectionTenant) error {
	if tenant.Ios.Enabled {
		var err error
		var authKey *ecdsa.PrivateKey
		var certificateKey tls.Certificate
		var ext string

		if tenant.Ios.KeyPath != "" {
			ext = filepath.Ext(tenant.Ios.KeyPath)

			switch ext {
			case dotP12:
				certificateKey, err = certificate.FromP12File(tenant.Ios.KeyPath, tenant.Ios.Password)
			case dotPEM:
				certificateKey, err = certificate.FromPemFile(tenant.Ios.KeyPath, tenant.Ios.Password)
			case dotP8:
				authKey, err = token.AuthKeyFromFile(tenant.Ios.KeyPath)
			default:
				err = errors.New("wrong certificate key extension")
			}

			if err != nil {
				logx.LogError.Error("Cert Error:", err.Error())

				return err
			}
		} else if tenant.Ios.KeyBase64 != "" {
			ext = "." + tenant.Ios.KeyType
			key, err := base64.StdEncoding.DecodeString(tenant.Ios.KeyBase64)
			if err != nil {
				logx.LogError.Error("base64 decode error:", err.Error())

				return err
			}
			switch ext {
			case dotP12:
				certificateKey, err = certificate.FromP12Bytes(key, tenant.Ios.Password)
			case dotPEM:
				certificateKey, err = certificate.FromPemBytes(key, tenant.Ios.Password)
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
			if tenant.Ios.KeyID == "" || tenant.Ios.TeamID == "" {
				msg := "you should provide ios.KeyID and ios.TeamID for p8 token"
				logx.LogError.Error(msg)
				return errors.New(msg)
			}
			jwt := &token.Token{
				AuthKey: authKey,
				// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
				KeyID: tenant.Ios.KeyID,
				// TeamID from developer account (View Account -> Membership)
				TeamID: tenant.Ios.TeamID,
			}

			ApnsClients[tenantId], err = newApnsTokenClient(jwt, tenant.Ios.Production)
		} else {
			ApnsClients[tenantId], err = newApnsClient(certificateKey, tenant.Ios.Production)
		}

		if h2Transport, ok := ApnsClients[tenantId].HTTPClient.Transport.(*http2.Transport); ok {
			configureHTTP2ConnHealthCheck(h2Transport)
		}

		if err != nil {
			logx.LogError.Error("Transport Error:", err.Error())

			return err
		}

		doOnce.Do(func() {
			MaxConcurrentIOSPushes[tenantId] = make(chan struct{}, tenant.Ios.MaxConcurrentPushes)
		})
	}

	return nil
}

func newApnsClient(certificate tls.Certificate, isProduction bool) (*apns2.Client, error) {
	var client *apns2.Client

	if isProduction {
		client = apns2.NewClient(certificate).Production()
	} else {
		client = apns2.NewClient(certificate).Development()
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

func newApnsTokenClient(token *token.Token, isProduction bool) (*apns2.Client, error) {
	var client *apns2.Client

	if isProduction {
		client = apns2.NewTokenClient(token).Production()
	} else {
		client = apns2.NewTokenClient(token).Development()
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

	if req.MutableContent {
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
		client = ApnsClients[req.TenantId].Production()
	case req.Development:
		client = ApnsClients[req.TenantId].Development()
	default:
		if cfg.Tenants[req.TenantId].Ios.Production {
			client = ApnsClients[req.TenantId].Production()
		} else {
			client = ApnsClients[req.TenantId].Development()
		}
	}

	return
}

// PushToIOS provide send notification to APNs server.
func PushToIOS(req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for iOS")
	if req.TenantId == "" {
		logx.LogError.Error("missing tenant id for Android notification")
		return
	}

	var (
		retryCount = 0
		maxRetry   = cfg.Tenants[req.TenantId].Ios.MaxRetry
	)

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	resp = &ResponsePush{}

Retry:
	var newTokens []string

	notification := GetIOSNotification(req)
	client := getApnsClient(cfg, req)

	var wg sync.WaitGroup
	for _, iosToken := range req.Tokens {
		// occupy push slot
		MaxConcurrentIOSPushes[req.TenantId] <- struct{}{}
		wg.Add(1)
		go func(notification apns2.Notification, token string) {
			notification.DeviceToken = iosToken

			// send ios notification
			res, err := client.Push(&notification)
			if err != nil || (res != nil && res.StatusCode != http.StatusOK) {
				if err == nil {
					// error message:
					// ref: https://github.com/sideshow/apns2/blob/master/response.go#L14-L65
					err = errors.New(res.Reason)
				}

				// apns server error
				errLog := logPush(cfg, core.FailedPush, iosToken, req, err)
				resp.Logs = append(resp.Logs, errLog)

				status.StatStorage.AddIosError(1)
				// We should retry only "retryable" statuses. More info about response:
				// See https://apple.co/3AdNane (Handling Notification Responses from APNs)
				if res != nil && res.StatusCode >= http.StatusInternalServerError {
					newTokens = append(newTokens, iosToken)
				}
			}

			if res != nil && res.Sent() {
				logPush(cfg, core.SucceededPush, iosToken, req, nil)
				status.StatStorage.AddIosSuccess(1)
			}

			// free push slot
			<-MaxConcurrentIOSPushes[req.TenantId]
			wg.Done()
		}(*notification, iosToken)
	}

	wg.Wait()

	if len(newTokens) > 0 && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return resp, nil
}
