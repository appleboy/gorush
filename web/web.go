package web

import (
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"strings"
	"strconv"
	"errors"
	"regexp"
	"io/ioutil"
	"github.com/martijnc/gowebpush/ece"
	"github.com/martijnc/gowebpush/webpush"
)

var (
	HTTPClientTimeout = 30 * time.Second
)

type Client struct {
	HTTPClient    *http.Client
}

type Response struct {
	StatusCode   int
	Body         string
}

type Subscription struct {
	Endpoint  string       `json:"endpoint,omitempty"`
	Key       string       `json:"key,omitempty"`
	Auth      string       `json:"auth,omitempty"`
}

type Notification struct {
	Subscription  *Subscription                 `json:"subscription,omitempty"`
	Payload       *map[string]interface{}       `json:"payload,omitempty"`
	TimeToLive    *uint                         `json:"time_to_live,omitempty"`
}

type Browser struct {
	Name      string
	ReDetect  regexp.Regexp
	ReError   regexp.Regexp
}

var Browsers = [...]Browser{
	Browser{"Chrome", *regexp.MustCompile("https://android.googleapis.com/gcm/send/"), *regexp.MustCompile("<TITLE>(.*)</TITLE>")},
	Browser{"Firefox", *regexp.MustCompile("https://updates.push.services.mozilla.com/wpush"), *regexp.MustCompile("\\\"error\\\":\\s\\\"([^\"]*)\\\"")},
}

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout:   HTTPClientTimeout,
		},
	}
}

func (c *Client) Push(n *Notification, apiKey string) (*Response, error) {
	jsonBuffer, _ := json.Marshal(n.Payload)
	var timeToLive uint
	if n.TimeToLive != nil {
		timeToLive = *n.TimeToLive
	} else {
		timeToLive = 2419200 
	}

	var authKey, p256dhKey []byte
	var errAuth, errKey error
	authKey, errAuth = base64.RawURLEncoding.DecodeString(n.Subscription.Auth)
	if errAuth != nil {
		authKey, errAuth = base64.URLEncoding.DecodeString(n.Subscription.Auth)
		if errAuth != nil {
			return nil, errAuth
		}
	}

	p256dhKey, errKey = base64.RawURLEncoding.DecodeString(n.Subscription.Key)
	if errKey != nil {
		p256dhKey, errKey = base64.URLEncoding.DecodeString(n.Subscription.Key)
		if errKey != nil {
			return nil, errKey
		}
	}
	var sp, rp webpush.KeyPair
	sp.GenerateKeys()
	err2 := rp.SetPublicKey(p256dhKey)
	if err2 != nil {
		return nil, err2
	}
	// Calculate the shared secret from the key-pairs (IKM).
	secret := webpush.CalculateSecret(&sp, &rp)

	var keys ece.EncryptionKeys
	encryptionContext := ece.BuildDHContext(rp.PublicKey, sp.PublicKey)
	keys.SetPreSharedAuthSecret(authKey)

	// Derive the encryption key and nonce from the input keying material.
	keys.CreateEncryptionKeys(secret, encryptionContext)

	// Encrypt the plaintext
	ciphertext, _ := ece.Encrypt(jsonBuffer, &keys, 25)

	// Create the headers
	var eh ece.EncryptionHeader
	eh.SetSalt(keys.GetSalt())
	eh.SetRecordSize(len(ciphertext))

	var ckh ece.CryptoKeyHeader
	ckh.SetDHKey(sp.PublicKey)

	var pushResponse Response

	// Create the ECE request.
	r := ece.CreateRequest(*c.HTTPClient, n.Subscription.Endpoint, ciphertext, &ckh, &eh, int(timeToLive))
	if strings.Contains(n.Subscription.Endpoint, "https://android.googleapis.com/gcm/send") {
		r.Header.Add("Authorization", "key=" + apiKey)
	}
	response, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	fmt.Println("Status Code: " + strconv.Itoa(response.StatusCode))
	pushResponse.StatusCode = response.StatusCode
	pushResponse.Body = string(body)
	if response.StatusCode != 201 {
		return &pushResponse, errors.New("Push endpoint returned incorrect status code: " + strconv.Itoa(response.StatusCode))
	}

	return &pushResponse, nil
}