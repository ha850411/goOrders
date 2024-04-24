package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"goOrders/database"
	"goOrders/models"
	"goOrders/service"
	"log"
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

	redis, err := service.GetRedisClient()
	if err != "" {
		log.Println(err)
	}
	token, _ := c.Cookie("loginToken")

	jsonData, _ := redis.Get(context.Background(), token).Result()
	user := models.Users{}
	_ = json.Unmarshal([]byte(jsonData), &user)

	output["username"] = user.Username
	output["active"] = active
	output["staticFreshFlag"] = time.Now().Unix()
	return output
}

// 首頁
func Index(c *gin.Context) {
	c.String(http.StatusOK, "test")
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
