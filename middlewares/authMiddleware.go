package middlewares

import (
	"ecommerce-rbac/logger"
	"ecommerce-rbac/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug("In AuthMiddleware function")
		rawHeader := c.GetHeader("Authorization")
		if rawHeader == "" || !strings.HasPrefix(rawHeader, "Bearer") {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		jwtToken := strings.Trim(strings.Replace(rawHeader, "Bearer", "", 1), " ")
		if jwtToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		username, err := service.ValidateToken(jwtToken, role)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("role", role)
		c.Next()
	}
}

func AuthMiddlewareWithMultipleRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug("In AuthMiddleware function")
		rawHeader := c.GetHeader("Authorization")
		if rawHeader == "" || !strings.HasPrefix(rawHeader, "Bearer") {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		jwtToken := strings.TrimSpace(strings.Replace(rawHeader, "Bearer", "", 1))
		if jwtToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		userClaims, err := service.ValidateTokenWithMultipleRole(jwtToken, roles) // Pass multiple roles
		// logger.Debug("this is the username we are gettitng ", username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("username", userClaims.Username)
		c.Set("role", userClaims.Role) // Store roles in context if needed
		c.Next()
	}
}
