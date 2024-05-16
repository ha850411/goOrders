package controllers

import (
	"database/sql"
	"goOrders/database"
	"goOrders/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func init() {
	db = database.DbConnect()
}

func GetCommonOutput(c *gin.Context, active string) map[string]interface{} {
	output := make(map[string]interface{})

	userInfo, _ := c.Get("userInfo")
	output["username"] = userInfo.(models.Users).Username
	output["active"] = active
	output["staticFreshFlag"] = time.Now().Unix()
	return output
}

// 自動證書驗證
func AutoCert(c *gin.Context) {
	filename := c.Param("files")
	b, err := os.ReadFile("./.well-known/acme-challenge/" + filename)
	if err != nil {
		NotFound(c)
	}
	c.String(http.StatusOK, string(b))
}

// 404 page
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", nil)
}
