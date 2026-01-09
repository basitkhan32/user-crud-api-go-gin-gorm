package controllers

import (
	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ReadUserController(c *gin.Context) {
	var users []models.User
	result := connections.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func ReadUserIDController(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := connections.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User with id " + string(id) + " fetched successfully",
		"user":    user,
	})
}
