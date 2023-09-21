package routes

import (
	controllers "github.com/aryamanchandra/supplify/controllers"
	middleware "github.com/aryamanchandra/supplify/middleware"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.POST("/newproduct", controllers.NewProduct)
	incomingRoutes.GET("/getsupplychain", controllers.GetSupplychain)
	incomingRoutes.POST("/addblock", controllers.WriteBlock)
}
