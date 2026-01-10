package auth

import (
	"github.com/basitkhan32/crud-api-go-gin/db"
	"github.com/basitkhan32/crud-api-go-gin/modules/user"
	"github.com/basitkhan32/crud-api-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	uBody := RegisterInput{}

	err := c.ShouldBindJSON(&uBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser := user.User{
		Name:     uBody.Name,
		Email:    uBody.Email,
		Age:      uBody.Age,
		Password: uBody.Password,
	}

	result := db.DB.Create(&newUser)
	if utils.ErrServerCrash(c, result.Error) {
		return
	}

	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"user":    newUser,
	})
}

func Login(c *gin.Context) {
	uBody := LoginInput{}

	err := c.ShouldBindJSON(&uBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var foundUser user.User
	result := db.DB.Where("email = ? AND password = ?", uBody.Email, uBody.Password).First(&foundUser)
	if utils.ErrBadRequest(c, result.Error) {
		return
	}

	c.JSON(200, gin.H{
		"message": "User logged in successfully",
		"user":    foundUser,
	})
}

func UpdatePassword(c *gin.Context) {
	uBody := UpdatePasswordInput{}

	err := c.ShouldBindJSON(&uBody)
	if utils.ErrBadRequest(c, err) {
		return
	}

	var foundUser user.User
	result := db.DB.Where("email = ? AND password = ?", uBody.Email, uBody.OldPassword).First(&foundUser)
	if utils.ErrBadRequest(c, result.Error) {
		return
	}

	// Update the password
	foundUser.Password = uBody.NewPassword
	updateResult := db.DB.Save(&foundUser)
	if utils.ErrServerCrash(c, updateResult.Error) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Password updated successfully",
		"user":    foundUser,
	})
}
