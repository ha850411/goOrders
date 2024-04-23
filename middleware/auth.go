package middleware

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 後台中間件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginToken, err := c.Cookie("loginToken")
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}
		username, _ := base64.RawURLEncoding.DecodeString(loginToken)
		c.Set("username", string(username))
		c.Next()
	}
}
