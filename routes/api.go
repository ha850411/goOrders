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

		productRouter := apiGroup.Group("product")
		{
			productRouter.GET("", api.GetProducts)
			productRouter.PUT("", api.UpdateProduct)
			productRouter.POST("", api.CreateProduct)
			productRouter.DELETE("/:id", api.DeleteProduct)
			productRouter.PATCH("/sort", api.SortProduct)
		}

		productTypeRouter := apiGroup.Group("productType")
		{
			productTypeRouter.GET("", api.GetProductType)
			productTypeRouter.GET("/", api.GetProductType)
		}

		settingRouter := apiGroup.Group("setting")
		{
			settingRouter.GET("/unsetLineNotify", api.UnsetLineNotify)
		}

		lineRouter := apiGroup.Group("line")
		{
			lineRouter.GET("/oauth", api.LineOauth)
			lineRouter.GET("/login", api.LineLogin)
		}
	}
}
