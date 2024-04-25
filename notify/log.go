package notify

import (
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
)

func logPush(cfg *config.ConfYaml, status, token string, req *PushNotification, err error) logx.LogPushEntry {
	return logx.LogPush(&logx.InputLog{
		ID:          req.ID,
		Status:      status,
		Token:       token,
		Message:     req.Message,
		Platform:    req.Platform,
		Error:       err,
		HideToken:   cfg.Log.HideToken,
		HideMessage: cfg.Log.HideMessages,
		Format:      cfg.Log.Format,
	})
}
