package user

import (
	"net/http"

	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// get data from request body

	newUser := User{}

	err := c.ShouldBindJSON(&newUser)
	utils.ErrBadRequest(c, err)

	result := connections.DB.Create(&newUser)
	utils.ErrServerCrash(c, result.Error)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}

func DeleteUser(c *gin.Context) {
	var deletedUser User
	id := c.Param("id")

	result := connections.DB.Delete(&deletedUser, id)
	utils.ErrServerCrash(c, result.Error)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully with id:" + id,
		"user":    deletedUser,
	})
}

func GetUser(c *gin.Context) {
	var users []User
	result := connections.DB.Find(&users)
	utils.ErrServerCrash(c, result.Error)

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func GetUserID(c *gin.Context) {
	var foundUser User
	id := c.Param("id")
	result := connections.DB.First(&foundUser, id)
	utils.ErrServerCrash(c, result.Error)

	c.JSON(http.StatusOK, gin.H{
		"message": "User with id " + string(id) + " fetched successfully",
		"user":    foundUser,
	})
}

func UpdateUser(c *gin.Context) {
	var existing_user User
	var update_user User
	id := c.Param("id")

	result := connections.DB.First(&existing_user, id)
	utils.ErrServerCrash(c, result.Error)

	// bind to update_user
	err := c.ShouldBindJSON(&update_user)
	utils.ErrBadRequest(c, err)

	update := connections.DB.Model(&existing_user).Updates(&update_user)
	utils.ErrServerCrash(c, update.Error)

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    existing_user,
	})

}
