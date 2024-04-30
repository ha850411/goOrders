package notify

import (
	"fmt"
	"goOrders/conf"
	"goOrders/database"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	REDIRECT_URI    = "http://localhost:8888/api/line/oauth"
	OAUTH_TOKEN_URL = "https://notify-bot.line.me/oauth/token"
	NOTICE_URL      = "https://notify-api.line.me/api/notify"
	CLIENT_ID       = conf.Settings.LineNotify.ClientID
	CLIENT_SECRET   = conf.Settings.LineNotify.ClientSecret
)

func init() {
	if conf.Settings.Common.ENV != "local" {
		REDIRECT_URI = "http://107.191.52.212/api/line/oauth"
	}
}

func Request(message string) {
	form := url.Values{}
	form.Add("message", message)

	sql := "SELECT linebot_token FROM users WHERE linebot_token IS NOT NULL"
	db := database.DbConnect()
	rows, err := db.Query(sql)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var token string
		rows.Scan(&token)
		bearerToken := fmt.Sprintf("Bearer %s", token)
		req, _ := http.NewRequest("POST", NOTICE_URL, strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", bearerToken)

		client := &http.Client{}
		resp, _ := client.Do(req)
		response, _ := ioutil.ReadAll(resp.Body)
		// var jsonResponse map[string]interface{}
		fmt.Printf("response: %v\n", string(response))
	}
}
