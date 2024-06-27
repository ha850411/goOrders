package controllers

import (
	"goOrders/service/line/login"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 登入頁
func IndexLogin(c *gin.Context) {
	// url 物件
	params := url.Values{}
	// 添加參數
	params.Add("response_type", "code")
	params.Add("client_id", login.CLIENT_ID)
	params.Add("redirect_uri", login.REDIRECT_URI)
	params.Add("state", "csrfToken")
	params.Add("scope", "profile openid email")
	// 產生 url
	baseUrl, _ := url.Parse(login.LOGIN_URL)
	baseUrl.RawQuery = params.Encode()

	c.HTML(http.StatusOK, "indexLogin", gin.H{
		"lineUrl": baseUrl.String(),
	})
}

// 首頁
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}
