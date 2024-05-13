package routes

import (
	"goOrders/controllers"
	"goOrders/middleware"

	"github.com/gin-gonic/gin"
)

func GetWebRouters(r *gin.Engine) {

	// ========= 後台 ===========
	adminGroup := r.Group("/admin")
	adminGroup.GET("/login", controllers.Login)   // 登入頁
	adminGroup.POST("/auth", controllers.Auth)    // 登入驗證
	adminGroup.GET("/logout", controllers.Logout) // 登出頁

	adminGroup.GET("/", middleware.Auth(), controllers.ProductManager)                 // 商品管理
	adminGroup.GET("/product", middleware.Auth(), controllers.ProductManager)          // 商品管理
	adminGroup.GET("/product-type", middleware.Auth(), controllers.ProductTypeManager) // 商品類別管理

	adminGroup.GET("/setting", middleware.Auth(), controllers.SettingManager) // 設定

	// ========= 前台 ===========
	// r.GET("/", middleware.CsrfHandler(), controllers.Index)

	// auto cert
	r.GET("/.well-known/acme-challenge/*files", controllers.AutoCert)
}
