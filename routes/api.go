package routes

import (
	"goOrders/controllers/api"
	"goOrders/middleware"

	"github.com/gin-gonic/gin"
)

func GetApiRouters(r *gin.Engine) {
	apiGroup := r.Group("api")
	{
		apiGroup.Use(middleware.Auth())

		producttRouter := apiGroup.Group("product")
		{
			producttRouter.GET("", api.GetProducts)
			producttRouter.GET("/", api.GetProducts)
		}

		settingRouter := apiGroup.Group("setting")
		{
			settingRouter.GET("/unsetLineNotify", api.UnsetLineNotify)
		}

		lineRouter := apiGroup.Group("line")
		{
			lineRouter.GET("/oauth", api.LineOauth)
		}
	}
}
