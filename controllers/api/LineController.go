package api

import (
	"encoding/json"
	"goOrders/conf"
	"goOrders/database"
	"goOrders/middleware"
	"goOrders/models"
	"goOrders/service"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func LineOauth(c *gin.Context) {
	code := c.Query("code")

	// 呼叫 linebot api 取得 access_token
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", service.REDIRECT_URI)
	form.Add("client_id", conf.Settings.LineNotify.ClientID)
	form.Add("client_secret", conf.Settings.LineNotify.ClientSecret)

	req, _ := http.NewRequest("POST", service.OAUTH_TOKEN_URL, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&respBody)

	db := database.GormConnect()
	userInfo, _ := c.Get("userInfo")
	result := db.Model(&models.Users{}).Where("username = ?", userInfo.(models.Users).Username).Update("linebot_token", respBody["access_token"])
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// 重新寫入 redis
	middleware.SetUserRedis(c)

	c.Redirect(http.StatusMovedPermanently, "/admin/setting")
}
