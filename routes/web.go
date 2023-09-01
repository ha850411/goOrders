package routes

import (
	"goOrders/controllers"
	"goOrders/middleware"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func GetWebRouters(r *gin.Engine) {
	// 載入共用模板設定
	r.HTMLRender = createMyRender()
	// ========= 後台 ===========
	// adminGroup := r.Group("/admin")
	// adminGroup.GET("/login", controllers.Login)   // 登入頁
	// adminGroup.POST("/auth", controllers.Auth)    // 登入驗證
	// adminGroup.GET("/logout", controllers.Logout) // 登出頁

	// ========= 前台 ===========
	r.GET("/", middleware.CsrfHandler(), controllers.Index)

	// auto cert
	r.GET("/.well-known/acme-challenge/*files", controllers.AutoCert)
}

/*
* 共用模板設定
* @return multitemplate.Renderer
 */
func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// common := map[string]string{
	// 	"menu":         "views/layout/menu.html",
	// 	"header":       "views/layout/header.html",
	// 	"layout":       "views/layout/layout.html",
	// 	"front-layout": "views/front/front-layout.html",
	// 	"front-header": "views/front/front-header.html",
	// }
	// // 存貨分析
	// r.AddFromFiles("analysis", common["layout"], common["header"], common["menu"], "views/main/analysis.html")

	// === 404 page ===
	r.AddFromFiles("404", "views/main/404.html")
	return r
}
