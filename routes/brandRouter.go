package routes

import (
	controllers "github.com/aryamanchandra/supplify/controllers"
	middleware "github.com/aryamanchandra/supplify/middleware"
	"github.com/gin-gonic/gin"
)

func BrandRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.POST("/new", controllers.NewProduct)
	incomingRoutes.GET("/supplychain", controllers.GetSupplychain)

}
