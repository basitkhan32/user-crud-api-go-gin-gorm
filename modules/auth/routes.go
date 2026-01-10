package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.RouterGroup) {
	r := router.Group("auth")
	{
		r.POST("/register", Register)
		r.POST("/login", Login)
		r.POST("/update-password", UpdatePassword)
	}
}
