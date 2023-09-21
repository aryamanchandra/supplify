package routes

import (
	controllers "github.com/aryamanchandra/supplify/controllers"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/newproduct", controllers.NewProduct)
	incomingRoutes.GET("/getsupplychain", controllers.GetSupplychain)
	incomingRoutes.POST("/addblock", controllers.WriteBlock)
}
