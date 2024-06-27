package routes

import "github.com/gin-contrib/multitemplate"

var admin = map[string]string{
	"menu":   "views/admin/menu.html",
	"layout": "views/admin/layout.html",
	"header": "views/admin/header.html",
}

var index = map[string]string{
	"layout": "views/index/layout.html",
}

/*
* 共用模板設定
* @return multitemplate.Renderer
 */
func createMyRender() multitemplate.Renderer {

	render := multitemplate.NewRenderer()

	// 取得後台模版設定
	getAdminRouterSetting(render)

	// 取得首頁模板設定
	getIndexRouterSetting(render)

	// 404 page
	render.AddFromFiles("404", "views/main/404.html")

	return render
}

func getAdminRouterSetting(render multitemplate.Renderer) {
	// 登入頁
	render.AddFromFiles("login", admin["layout"], "views/admin/login.html")
	// 商品頁
	render.AddFromFiles("productList", admin["layout"], admin["header"], admin["menu"], "views/admin/product.html")
	// 商品類別
	render.AddFromFiles("productTypeList", admin["layout"], admin["header"], admin["menu"], "views/admin/product-type.html")
	// 桌號管理
	render.AddFromFiles("deskList", admin["layout"], admin["header"], admin["menu"], "views/admin/desk.html")
	// 設定頁
	render.AddFromFiles("setting", admin["layout"], admin["header"], admin["menu"], "views/admin/setting.html")
}

func getIndexRouterSetting(render multitemplate.Renderer) {
	// 首頁
	render.AddFromFiles("index", index["layout"], "views/index/index.html")
	// 登入頁
	render.AddFromFiles("indexLogin", index["layout"], "views/index/login.html")
}
