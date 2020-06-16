package gorush

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"encoding/json"

	"github.com/msalihkarakasli/go-hms-push/push/model"
	"github.com/msalihkarakasli/go-hms-push/push/core"
	"github.com/msalihkarakasli/go-hms-push/push/config"
)

var (
	pushError   error
	pushClient  *core.HMSClient
	once        sync.Once
)

// GetPushClient use for create HMS Push
func GetPushClient(conf *config.Config) (*core.HMSClient, error){
	once.Do(func() {
		client, err := core.NewHttpClient(conf)
		if err != nil {
			fmt.Printf("Failed to new common client! Error is %s\n", err.Error())
			panic(err)
		}
		pushClient = client
		pushError = err
	})

	return pushClient, pushError
}

// InitHMSClient use for initialize HMS Client.
func InitHMSClient(apiKey string, appId string) (*core.HMSClient, error) {

	if apiKey == "" {
		return nil, errors.New("Missing Huawei API Key")
	}

	if appId == "" {
		return nil, errors.New("Missing Huawei App Id")
	}

	var conf = &config.Config{
		AppId:     appId,
		AppSecret: apiKey,
		AuthUrl: "https://login.cloud.huawei.com/oauth2/v2/token",
		PushUrl: "https://api.push.hicloud.com",
	}

	if apiKey != PushConf.Huawei.APIKey || appId != PushConf.Huawei.APPId {
		return GetPushClient(conf)
	}

	if HMSClient == nil {
		return GetPushClient(conf)
	}

	return HMSClient, nil
}

// GetHuaweiNotification use for define HMS notification.
// HTTP Connection Server Reference for HMS
// https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi
func GetHuaweiNotification(req PushNotification) (*model.MessageRequest, error) {
	
	msgRequest := model.NewNotificationMsgRequest()

	msgRequest.Message.Android = model.GetDefaultAndroid()

	if len(req.Tokens) > 0 {
		msgRequest.Message.Token = req.Tokens
	}

	if len(req.Topic) > 0 {
		msgRequest.Message.Topic = req.Topic;
	}

	if len(req.Condition) > 0 {
		msgRequest.Message.Condition = req.Condition
	}

	if req.Priority == "high" {
		msgRequest.Message.Android.Urgency = "HIGH"
	}

	//if req.HuaweiCollapseKey != nil {
		msgRequest.Message.Android.CollapseKey = req.HuaweiCollapseKey
	//}

	if len(req.Category) > 0 {
		msgRequest.Message.Android.Category = req.Category
	}

	if len(req.HuaweiTTL) > 0 {
		msgRequest.Message.Android.TTL = req.HuaweiTTL
	}

	if len(req.BiTag) > 0 {
		msgRequest.Message.Android.BiTag = req.BiTag
	}

	//if req.FastAppTarget != nil {
		msgRequest.Message.Android.FastAppTarget = req.FastAppTarget
	//}

	//Add data fields
	if len(req.HuaweiData) > 0 {
	  	msgRequest.Message.Data = req.HuaweiData
	} else {
		//Notification Message
		msgRequest.Message.Android.Notification = model.GetDefaultAndroidNotification()

		n := msgRequest.Message.Android.Notification
		isNotificationSet := false
		
		if req.HuaweiNotification != nil {
		 	isNotificationSet = true
		 	n = req.HuaweiNotification

		 	if n.ClickAction == nil {
		 		n.ClickAction = model.GetDefaultClickAction()
		 	}
		}

		if len(req.Message) > 0 {
		 	isNotificationSet = true
		 	n.Body = req.Message
		}

		if len(req.Title) > 0 {
		 	isNotificationSet = true
		 	n.Title = req.Title
		}

		if len(req.Image) > 0 {
		 	isNotificationSet = true
		 	n.Image = req.Image
		}

		if v, ok := req.Sound.(string); ok && len(v) > 0 {
		 	isNotificationSet = true
		 	n.Sound = v
		} else {
			n.DefaultSound = true;
		}


		if isNotificationSet {
		 	msgRequest.Message.Android.Notification = n
		 }
	}

	// if len(req.Apns) > 0 {
	// 	notification.Apns = req.Apns
	// }

	b, err := json.Marshal(msgRequest)
	if err != nil {
		fmt.Printf("Failed to marshal the default message! Error is %s\n", err.Error())
		return nil, err
	}

	fmt.Printf("Default message is %s\n", string(b))
	return msgRequest, nil
}

// PushToHuawei provide send notification to Android server.
func PushToHuawei(req PushNotification) bool {
	LogAccess.Debug("Start push notification for Huawei")

	var (
		client     *core.HMSClient
		retryCount = 0
		maxRetry   = PushConf.Huawei.MaxRetry
	)

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

	notification, err := GetHuaweiNotification(req)

	client, err = InitHMSClient(PushConf.Huawei.APIKey, PushConf.Huawei.APPId)

	if err != nil {
		// HMS server error
		LogError.Error("HMS server error: " + err.Error())
		return false
	}

	res, err := client.SendMessage(context.Background(), notification)
	if err != nil {
		// Send Message error
		LogError.Error("HMS server send message error: " + err.Error())
		return false
	}

	fmt.Println(res.Code);
	fmt.Println(res.Msg);
	fmt.Println(res.RequestId);


	// Huawei Push Send API does not support exact results for each token
	if res.Code=="80000000" {
		StatStorage.AddHuaweiSuccess(int64(1))
		LogAccess.Debug(fmt.Sprintf("Huwaei Send Notification is completed successfully!"))
	} else {
		isError = true
		StatStorage.AddHuaweiError(int64(1))
		LogAccess.Debug(fmt.Sprintf("Huwaei Send Notification is failed!"))
	}

	if isError && retryCount < maxRetry {
	 	retryCount++

	 	// resend all tokens
	 	goto Retry
	}

	return isError
}
