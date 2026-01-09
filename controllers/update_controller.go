package controllers

import (
	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func UpdateUserController(c *gin.Context) {
	var existing_user models.User
	var update_user models.User
	id := c.Param("id")

	result := connections.DB.First(&existing_user, id)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// bind to update_user
	err := c.ShouldBindJSON(&update_user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	update := connections.DB.Model(&existing_user).Updates(&update_user)
	if update.Error != nil {
		c.JSON(500, gin.H{
			"error": update.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User updated successfully",
		"user":    update_user,
	})

}
