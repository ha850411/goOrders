package service

import (
	"goOrders/conf"
)

var (
	REDIRECT_URI    = ""
	OAUTH_TOKEN_URL = "https://notify-bot.line.me/oauth/token"
	NOTICE_URL      = "https://notify-api.line.me/api/notify"
	CLIENT_ID       = conf.Settings.LineNotify.ClientID
	CLIENT_SECRET   = conf.Settings.LineNotify.ClientSecret
)

func init() {
	if conf.Settings.Common.ENV == "local" {
		REDIRECT_URI = "http://localhost:8888/linebot/notify"
	} else {
		REDIRECT_URI = "http://107.191.52.212/linebot/notify"
	}
}
