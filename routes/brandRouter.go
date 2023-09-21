package routes

import (
	controllers "github.com/aryamanchandra/supplify/controllers"
	"github.com/gin-gonic/gin"
)

func BrandRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/new", controllers.NewProduct)
	incomingRoutes.GET("/supplychain", controllers.GetSupplychain)

}
