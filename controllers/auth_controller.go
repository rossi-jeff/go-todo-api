package controllers

import (
	"go-todo-api/dto"
	"go-todo-api/initializers"
	"go-todo-api/models"
	"go-todo-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRegister(c *gin.Context) {
	body := dto.RegisterCredentials{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hashed, err2 = utilities.HashPassword(body.PassWord)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	user := models.User{
		UserName:       body.UserName,
		Email:          body.Email,
		PassWordDigest: hashed,
	}

	initializers.DB.Save(&user)

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Save User"})
		return
	}

	// do not display PassWordDigest
	response := models.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		IsDeleted: user.IsDeleted,
		Random:    user.Random,
	}

	c.JSON(http.StatusCreated, response)
}

func AuthLogin(c *gin.Context) {
	body := dto.LoginCredentials{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	initializers.DB.Where("UserName = ?", body.UserName).First(&user)

	if !utilities.PasswordsMatch(user.PassWordDigest, body.PassWord) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	response := dto.LoginResponse{
		UserName: body.UserName,
	}
	c.JSON(http.StatusOK, response)
}
