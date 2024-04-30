package controllers

import (
	"fmt"
	"goOrders/models"
	"goOrders/service/line/notify"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SettingManager(c *gin.Context) {
	output := GetCommonOutput(c, "setting")

	userInfo, _ := c.Get("userInfo")

	output["bindLineNotify"] = userInfo.(models.Users).LinebotToken != ""

	output["lineUrl"] = fmt.Sprintf("https://notify-bot.line.me/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=notify&state=csrfToken", notify.CLIENT_ID, notify.REDIRECT_URI)

	c.HTML(http.StatusOK, "setting", output)
}
