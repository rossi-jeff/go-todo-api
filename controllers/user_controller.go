package controllers

import (
	"go-todo-api/dto"
	"go-todo-api/initializers"
	"go-todo-api/models"
	"go-todo-api/utilities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary User List
// @Description Get a list of users
// @Tags User
// @Accept       json
// @Produce      json
// @Success 200 {array} models.UserResponse
// @Router /user [get]
func UserList(c *gin.Context) {
	users := []models.UserResponse{} // no pass_word_digest

	initializers.DB.Model(&models.User{}).Select(
		"id", "user_name", "email", "random", "created_at", "updated_at", "is_deleted",
	).Find(&users)

	c.JSON(http.StatusOK, users)
}

// @Summary Current User
// @Description Get the currently logged in user
// @Tags User
// @Accept       json
// @Produce      json
// @Success 200 {object} models.UserResponse
// @Failure      401  {object} dto.ErrorResponse
// @Router /user/current [get]
// @Security Bearer
func UserCurrent(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := models.UserResponse{}

	initializers.DB.Model(&models.User{}).Where("id = ?", userId).Select(
		"id", "user_name", "email", "random", "created_at", "updated_at", "is_deleted",
	).Find(&user)

	c.JSON(http.StatusOK, user)
}

// @Summary Update User
// @Description Update a user
// @Tags User
// @Accept       json
// @Produce      json
// @Param UserID path int true "User ID"
// @Param	data	body	dto.UserAttrs true "User Attributes"
// @Success 200 {object} models.UserResponse
// @Failure      401  {object} dto.ErrorResponse
// @Failure      400  {object} dto.ErrorResponse
// @Router /user/{UserID} [patch]
// @Security Bearer
func UserUpdate(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	idStr := c.Param("id")
	id, err2 := strconv.ParseUint(idStr, 10, 8)
	if userId == 0 || err2 != nil || userId != uint(id) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	body := dto.UserAttrs{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.UserResponse{}

	initializers.DB.Model(&models.User{}).Select(
		"id", "user_name", "email", "random", "created_at", "updated_at", "is_deleted",
	).First(&user, idStr)
	user.Email = body.Email
	user.UserName = body.UserName

	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

// @Summary Delete User
// @Description Delete a user
// @Tags User
// @Accept       json
// @Produce      json
// @Param UserID path int true "User ID"
// @Success 204
// @Failure      401  {object} dto.ErrorResponse
// @Router /user/{UserID} [delete]
// @Security Bearer
func UserDelete(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	idStr := c.Param("id")
	id, err2 := strconv.ParseUint(idStr, 10, 8)
	if userId == 0 || err2 != nil || userId != uint(id) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := models.User{}

	initializers.DB.Model(&models.User{}).First(&user, idStr)
	user.IsDeleted = true

	initializers.DB.Save(&user)

	c.Writer.WriteHeader(http.StatusNoContent)
}

// @Summary Change PassWord
// @Description Change a PassWord for a user
// @Tags User
// @Accept       json
// @Produce      json
// @Param	data	body	dto.ChangePW true "PassWord Change Attributes"
// @Success 200 {object} models.UserResponse
// @Failure      401  {object} dto.ErrorResponse
// @Failure      400  {object} dto.ErrorResponse
// @Router /user/change [patch]
// @Security Bearer
func UserChangePassWord(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	body := dto.ChangePW{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if body.NewPassWord == "" || body.OldPassWord == "" || body.NewPassWord != body.Confirmation {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := models.User{}
	initializers.DB.First(&user, userId)

	if !utilities.PasswordsMatch(user.PassWordDigest, body.OldPassWord) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var hashed, _ = utilities.HashPassword(body.NewPassWord)
	user.PassWordDigest = hashed

	initializers.DB.Save(&user)

	response := models.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		IsDeleted: user.IsDeleted,
		Random:    user.Random,
	}

	c.JSON(http.StatusOK, response)
}
