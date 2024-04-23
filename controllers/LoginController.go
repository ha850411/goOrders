package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"goOrders/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登入頁
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{})
}

// 登入驗證
func Auth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 將密碼做 md5
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])
	fmt.Printf("password: %v\n", password)

	// db connect
	db := database.DbConnect()
	var count int
	sql := fmt.Sprintf("SELECT count(*) FROM users WHERE username='%s' AND password='%s'", username, password)
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	// base64
	token := base64.RawURLEncoding.EncodeToString([]byte(username))
	if count >= 1 {
		// 記住我
		remember := c.PostForm("remember")
		expire := 3600
		if remember == "Y" {
			expire = 86400 * 30
		}
		// 設定登入 token
		c.SetCookie("loginToken", token, expire, "/", "", false, false)
		// 轉導到後台首頁
		c.Redirect(http.StatusMovedPermanently, "/admin")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "帳號密碼錯誤, 請重新登入",
		})
	}
}

// 登出
func Logout(c *gin.Context) {
	c.SetCookie("loginToken", "", -1, "/", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/admin/login")
}
