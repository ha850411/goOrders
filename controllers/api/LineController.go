package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"goOrders/conf"
	"goOrders/database"
	"goOrders/middleware"
	"goOrders/models"
	"goOrders/service"
	"goOrders/service/line/login"
	"goOrders/service/line/notify"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LineOauth(c *gin.Context) {
	code := c.Query("code")

	// 呼叫 linebot api 取得 access_token
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", notify.REDIRECT_URI)
	form.Add("client_id", conf.Settings.LineNotify.ClientID)
	form.Add("client_secret", conf.Settings.LineNotify.ClientSecret)

	req, _ := http.NewRequest("POST", notify.OAUTH_TOKEN_URL, strings.NewReader(form.Encode()))
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

func LineLogin(c *gin.Context) {
	code := c.Query("code")

	// 呼叫 line oauth api 取得 access_token
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", login.REDIRECT_URI)
	form.Add("client_id", login.CLIENT_ID)
	form.Add("client_secret", login.CLIENT_SECRET)

	req, _ := http.NewRequest("POST", login.OAUTH_TOKEN_URL, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&respBody)
	accessToken := respBody["access_token"].(string)

	// 使用 accessToken 取得 profile
	req, _ = http.NewRequest("GET", login.PROFILE_URL, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client = &http.Client{}
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&respBody)

	// 分割 JWT 字串
	jwtToken := respBody["id_token"].(string)
	parts := strings.Split(jwtToken, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT token format")
		return
	}

	// 解碼並解析 payload
	payload, err := decodeSegment(parts[1])
	if err != nil {
		fmt.Printf("Error decoding segment: %v\n", err)
		return
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		fmt.Printf("Error unmarshalling payload: %v\n", err)
		return
	}

	// 幫使用者建立帳號, 寫入資料庫
	db := database.GormConnect()
	userInfo := &models.Users{}
	db.Where("username = ?", respBody["userId"].(string)).First(&userInfo)
	// 判斷使用者是否存在
	if userInfo.Username == "" {
		userInfo = &models.Users{
			Username: respBody["userId"].(string),
			Password: "",
			Role:     0,
			Name:     respBody["displayName"].(string),
			Email:    claims["email"].(string),
		}
		db.Create(userInfo)
	}

	// 寫入登入資訊
	uid := base64.URLEncoding.EncodeToString([]byte(userInfo.Username))
	c.SetCookie("uid", uid, 86400, "/", "", false, false)
	// 寫入 redis
	redisClient, _ := service.GetRedisClient()
	jsonData, _ := json.Marshal(userInfo)
	redisClient.Set(context.Background(), userInfo.Username, jsonData, time.Second*3600)

	c.Redirect(http.StatusFound, "/")
}

// decodeSegment 解碼 JWT 的一個段
func decodeSegment(segment string) ([]byte, error) {
	// 添加缺失的 '=' 符號以達到正確的 base64 編碼長度
	if l := len(segment) % 4; l > 0 {
		segment += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(segment)
}
