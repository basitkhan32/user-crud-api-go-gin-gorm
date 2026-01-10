package routes

import (
	"github.com/basitkhan32/crud-api-go-gin/modules/auth"
	"github.com/basitkhan32/crud-api-go-gin/modules/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	// Public routes - no middleware
	publicAPI := r.Group("/")
	{
		user.GetPublicRoutes(publicAPI)
		auth.AuthRoutes(publicAPI)
	}

	// Protected routes - with middleware
	protectedAPI := r.Group("/")
	protectedAPI.Use(user.IsXAPIProvided())
	{
		user.GetProtectedRoutes(protectedAPI)
	}
}
