package routes

import "github.com/gin-contrib/multitemplate"

/*
* 共用模板設定
* @return multitemplate.Renderer
 */
func createMyRender() multitemplate.Renderer {

	render := multitemplate.NewRenderer()

	admin := map[string]string{
		"menu":   "views/admin/menu.html",
		"layout": "views/admin/layout.html",
		"header": "views/admin/header.html",
	}
	// 登入頁
	render.AddFromFiles("login", admin["layout"], "views/admin/login.html")

	// 商品頁
	render.AddFromFiles("productList", admin["layout"], admin["header"], admin["menu"], "views/admin/product.html")

	// 設定頁
	render.AddFromFiles("setting", admin["layout"], admin["header"], admin["menu"], "views/admin/setting.html")

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
	render.AddFromFiles("404", "views/main/404.html")
	return render
}
