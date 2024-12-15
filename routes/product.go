package routes

import (
	"ecommerce-rbac/handlers"
	"ecommerce-rbac/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.RouterGroup) {
	// Web Builder route (accessible only by Sellers)
	sellerGroup := router.Group("/web-builder")
	sellerGroup.Use(middlewares.AuthMiddleware("seller"))
	sellerGroup.POST("/", handlers.CreateWebBuild)

	sellerProductGroup := router.Group("/products")
	sellerProductGroup.Use(middlewares.AuthMiddlewareWithMultipleRole("seller", "wholesaler"))
	sellerProductGroup.POST("/", handlers.CreateProduct)

	// wholesalerProductGroup := router.Group("/products")
	// wholesalerProductGroup.Use(middlewares.AuthMiddleware("wholesaler"))
	// wholesalerProductGroup.POST("/", handlers.CreateProduct)
}

