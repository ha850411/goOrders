package controllers

import (
	"encoding/json"
	"fmt"
	"goOrders/models"
	"goOrders/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SettingManager(c *gin.Context) {
	output := GetCommonOutput(c, "setting")

	var userInfo models.Users
	jsonStr, _ := c.Get("userInfo")
	fmt.Printf("userInfo: %v\n", userInfo)

	json.Unmarshal([]byte(jsonStr.(string)), &userInfo)

	output["bindLineNotify"] = userInfo.LinebotToken != ""

	output["lineUrl"] = fmt.Sprintf("https://notify-bot.line.me/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=notify&state=csrfToken", service.CLIENT_ID, service.REDIRECT_URI)

	c.HTML(http.StatusOK, "setting", output)
}
