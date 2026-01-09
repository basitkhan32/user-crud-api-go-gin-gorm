package user

import (
	"github.com/gin-gonic/gin"
)

func GetPublicRoutes(router *gin.RouterGroup) {
	u := router.Group("user")

	// Public route - no middleware
	u.POST("/create", CreateUser)
}

func GetProtectedRoutes(router *gin.RouterGroup) {
	u := router.Group("user")
	{
		u.POST("/", CreateUser)
		u.GET("/all", GetUser)
		u.GET("/:id", GetUserID)
		u.DELETE("/:id", DeleteUser)
		u.PUT("/:id", UpdateUser)
	}
}
