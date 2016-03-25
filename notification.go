package main

import (
	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	_ "github.com/google/go-gcm"
	"log"
)

type ExtendJSON struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

type RequestPushNotification struct {
	// Common
	Tokens   []string `json:"tokens" binding:"required"`
	Platform int      `json:"platform" binding:"required"`
	Message  string   `json:"message" binding:"required"`
	Priority string   `json:"priority,omitempty"`
	// Android
	CollapseKey    string `json:"collapse_key,omitempty"`
	DelayWhileIdle bool   `json:"delay_while_idle,omitempty"`
	TimeToLive     int    `json:"time_to_live,omitempty"`
	// iOS
	ApnsID string       `json:"apns_id,omitempty"`
	Topic  string       `json:"topic,omitempty"`
	Badge  int          `json:"badge,omitempty"`
	Sound  string       `json:"sound,omitempty"`
	Expiry int          `json:"expiry,omitempty"`
	Retry  int          `json:"retry,omitempty"`
	Extend []ExtendJSON `json:"extend,omitempty"`
	// meta
	IDs []uint64 `json:"seq_id,omitempty"`
}

func pushNotification(notification RequestPushNotification) bool {
	var (
		success bool
	)

	cert, err := certificate.FromPemFile("./key.pem", "")
	if err != nil {
		log.Println("Cert Error:", err)
	}

	apnsClient := apns.NewClient(cert).Development()

	switch notification.Platform {
	case PlatFormIos:
		success = pushNotificationIos(notification, apnsClient)
		if !success {
			apnsClient = nil
		}
	case PlatFormAndroid:
		success = pushNotificationAndroid(notification)
	}

	return success
}

func pushNotificationIos(req RequestPushNotification, client *apns.Client) bool {

	for _, token := range req.Tokens {
		notification := &apns.Notification{}
		notification.DeviceToken = token

		if len(req.ApnsID) > 0 {
			notification.ApnsID = req.ApnsID
		}

		if len(req.Topic) > 0 {
			notification.Topic = req.Topic
		}

		if len(req.Priority) > 0 && req.Priority == "low" {
			notification.Priority = apns.PriorityLow
		}

		payload := payload.NewPayload().Alert(req.Message)

		if req.Badge > 0 {
			payload.Badge(req.Badge)
		}

		if len(req.Sound) > 0 {
			payload.Sound(req.Sound)
		}

		if len(req.Extend) > 0 {
			for _, extend := range req.Extend {
				payload.Custom(extend.Key, extend.Value)
			}
		}

		notification.Payload = payload

		// send ios notification
		res, err := client.Push(notification)

		if err != nil {
			log.Println("There was an error", err)
			return false
		}

		if res.Sent() {
			log.Println("APNs ID:", res.ApnsID)
		}
	}

	client = nil

	return true
}

func pushNotificationAndroid(req RequestPushNotification) bool {

	return true
}
