package routes

import (
	"goOrders/controllers/api"

	"github.com/gin-gonic/gin"
)

func GetApiRouters(r *gin.Engine) {
	apiGroup := r.Group("api")
	{
		// apiGroup.Use(middleware.Auth())

		productRouter := apiGroup.Group("product")
		{
			productRouter.GET("", api.GetProducts)
			productRouter.PUT("", api.UpdateProduct)
			productRouter.POST("", api.CreateProduct)
			productRouter.DELETE("/:id", api.DeleteProduct)
			productRouter.PATCH("/sort", api.SortProduct)
		}

		productTypeRouter := apiGroup.Group("product-type")
		{
			productTypeRouter.GET("", api.GetProductType)
			productTypeRouter.PUT("", api.UpdateProductType)
			productTypeRouter.POST("", api.CreateProductType)
			productTypeRouter.DELETE("/:id", api.DeleteProductType)
		}

		deskRouter := apiGroup.Group("desk")
		{
			deskRouter.GET("", api.GetDesks)
			deskRouter.PUT("", api.UpdateDesk)
			deskRouter.POST("", api.CreateDesk)
			deskRouter.DELETE("/:id", api.DeleteDesk)
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
