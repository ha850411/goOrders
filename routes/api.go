package routes

import (
	"goOrders/controllers/api"

	"github.com/gin-gonic/gin"
)

func GetApiRouters(r *gin.Engine) {
	apiGroup := r.Group("api")
	{
		producttRouter := apiGroup.Group("product")
		{
			producttRouter.GET("", api.GetProducts)
			producttRouter.GET("/", api.GetProducts)
		}
	}
}
