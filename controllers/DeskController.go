package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeskManager(c *gin.Context) {
	output := GetCommonOutput(c, "desk")
	c.HTML(http.StatusOK, "deskList", output)
}
