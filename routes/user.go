package routes

import (
	"ecommerce-rbac/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRouters(router *gin.RouterGroup) {
	routeGroup := router.Group("/user")
	routeGroup.POST("/create", handlers.CreateUser)
	routeGroup.POST("/login", handlers.SignIn)
}
