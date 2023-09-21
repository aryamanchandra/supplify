package routes

import (
	controllers "github.com/aryamanchandra/supplify/controllers"
	"github.com/gin-gonic/gin"
)

func BuyerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/getsupplychain", controllers.GetSupplychain)
}
