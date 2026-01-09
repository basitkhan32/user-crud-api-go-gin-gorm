package controllers

import (
	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {
	// get data from request body

	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := connections.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
