package main

import (
	"os"

	routes "github.com/aryamanchandra/supplify/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.SellerRoutes(router)
	routes.BuyerRoutes(router)
	routes.BrandRoutes(router)

	router.Run(":" + port)
}
