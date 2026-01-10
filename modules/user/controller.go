package user

import (
	"net/http"

	"github.com/basitkhan32/crud-api-go-gin/db"
	"github.com/basitkhan32/crud-api-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// get data from request body

	newUser := User{}

	err := c.ShouldBindJSON(&newUser)
	if utils.ErrBadRequest(c, err) {
		return
	}

	result := db.DB.Create(&newUser)
	if utils.ErrServerCrash(c, result.Error) {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}

func DeleteUser(c *gin.Context) {
	var deletedUser User
	id := c.Param("id")

	result := db.DB.Delete(&deletedUser, id)
	if utils.ErrServerCrash(c, result.Error) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully with id:" + id,
		"user":    deletedUser,
	})
}

func GetUser(c *gin.Context) {
	var users []User
	result := db.DB.Find(&users)
	if utils.ErrServerCrash(c, result.Error) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func GetUserID(c *gin.Context) {
	var foundUser User
	id := c.Param("id")
	result := db.DB.First(&foundUser, id)
	if utils.ErrServerCrash(c, result.Error) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User with id " + string(id) + " fetched successfully",
		"user":    foundUser,
	})
}

func UpdateUser(c *gin.Context) {
	var existing_user User
	var update_user User
	id := c.Param("id")

	result := db.DB.First(&existing_user, id)
	if utils.ErrServerCrash(c, result.Error) {
		return
	}

	// bind to update_user
	err := c.ShouldBindJSON(&update_user)
	if utils.ErrBadRequest(c, err) {
		return
	}

	update := db.DB.Model(&existing_user).Updates(&update_user)
	if utils.ErrServerCrash(c, update.Error) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    existing_user,
	})

}
