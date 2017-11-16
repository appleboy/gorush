package gorush

import (
	"sync"
)

// InitWorkers for initialize all workers.
func InitWorkers(workerNum int64, queueNum int64) {
	LogAccess.Debug("worker number is ", workerNum, ", queue number is ", queueNum)
	QueueNotification = make(chan PushNotification, queueNum)
	for i := int64(0); i < workerNum; i++ {
		go startWorker()
	}
}

// SendNotification is send message to iOS or Android
func SendNotification(msg PushNotification) {
	switch msg.Platform {
	case PlatformIos:
		PushToIOS(msg)
	case PlatformAndroid:
		PushToAndroid(msg)
	case PlatformWeb:
		PushToWeb(msg)
	}
}

func startWorker() {
	for {
		notification := <-QueueNotification
		SendNotification(notification)
	}
}

// queueNotification add notification to queue list.
func queueNotification(req RequestPush) (int, []LogPushEntry) {
	var count int
	var doSync = PushConf.Core.Sync
	if (req.Sync != nil) {
		doSync = *req.Sync
	}
	wg := sync.WaitGroup{}
	newNotification := []PushNotification{}
	for _, notification := range req.Notifications {
		switch notification.Platform {
		case PlatformIos:
			if !PushConf.Ios.Enabled && !notification.Voip {
				continue
			}
			if !PushConf.Ios.VoipEnabled && notification.Voip {
				continue
			}
		case PlatformAndroid:
			if !PushConf.Android.Enabled {
				continue
			}
		case PlatformWeb:
			if !PushConf.Web.Enabled {
				continue
			}
		}
		notification.sync = doSync
		newNotification = append(newNotification, notification)
	}

	log := make([]LogPushEntry, 0, count)
	for _, notification := range newNotification {
		if doSync {
			notification.wg = &wg
			notification.log = &log
			notification.AddWaitCount()
		}
		QueueNotification <- notification
		switch notification.Platform {
		case PlatformWeb:
			count += len(notification.Subscriptions)
		default:
			count += len(notification.Tokens)
		}
		// Count topic message
		if notification.To != "" {
			count++
		}
	}

	if doSync {
		wg.Wait()
	}

	StatStorage.AddTotalCount(int64(count))

	return count, log
}
