package gorush

import (
	"fmt"
	"errors"

	"github.com/jaraxasoftware/gorush/web"
)

// InitWebClient use for initialize APNs Client.
func InitWebClient() error {
    if PushConf.Web.Enabled {
        //var err error
        WebClient = web.NewClient(PushConf.Web.APIKey)
    }

    return nil
}

func GetWebNotification(req PushNotification, subscription *Subscription) *web.Notification {
	notification := &web.Notification{
		Payload: &req.Data,
		Subscription: &web.Subscription{
			Endpoint: subscription.Endpoint, 
			Key: subscription.Key, 
			Auth: subscription.Auth,
		},
		TimeToLive: req.TimeToLive,
	}
	return notification
}

// PushToWeb provide send notification to Web server.
func PushToWeb(req PushNotification) bool {
	LogAccess.Debug("Start push notification for Web")
	var doSync = req.sync
	if doSync {
		defer req.WaitDone()
	}

	var retryCount = 0
	var maxRetry = PushConf.Web.MaxRetry

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err := CheckMessage(req)

	if err != nil {
		LogError.Error("request error: " + err.Error())
		return false
	}

Retry:
	var isError = false

	successCount := 0
	failureCount := 0

	for _, subscription := range req.Subscriptions {
		notification := GetWebNotification(req, &subscription)
		response, err := WebClient.Push(notification)
		if err != nil {
			failureCount++
			LogPush(FailedPush, subscription.Endpoint, req, err)
			fmt.Println(err)
			if doSync {
				var errorObj = errors.New(response.Body)
				req.AddLog(getLogPushEntry(FailedPush, subscription.Endpoint, req, errorObj))
			} 
		} else {
			successCount++
			LogPush(SucceededPush, subscription.Endpoint, req, nil)
		}
	}

	LogAccess.Debug(fmt.Sprintf("Web Success count: %d, Failure count: %d", successCount, failureCount))
	StatStorage.AddWebSuccess(int64(successCount))
	StatStorage.AddWebError(int64(failureCount))

	if isError && retryCount < maxRetry {
		retryCount++

		goto Retry
	}

	return isError  	
}
