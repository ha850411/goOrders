package login

import (
	"goOrders/conf"
)

var (
	REDIRECT_URI    = "http://localhost:8888/api/line/login"
	OAUTH_TOKEN_URL = "https://api.line.me/oauth2/v2.1/token"
	LOGIN_URL       = "https://access.line.me/oauth2/v2.1/authorize"
	PROFILE_URL     = "https://api.line.me/v2/profile"
	CLIENT_ID       = conf.Settings.LineLogin.ClientID
	CLIENT_SECRET   = conf.Settings.LineLogin.ClientSecret
)

func init() {
	if conf.Settings.Common.ENV != "local" {
		REDIRECT_URI = "http://107.191.52.212/api/line/login"
	}
}
