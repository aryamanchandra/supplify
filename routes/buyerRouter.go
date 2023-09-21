package routes

import (
	controllers "github.com/aryamanchandra/supplify/controllers"
	middleware "github.com/aryamanchandra/supplify/middleware"
	"github.com/gin-gonic/gin"
)

func BuyerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/getsupplychain", controllers.GetSupplychain)
}
