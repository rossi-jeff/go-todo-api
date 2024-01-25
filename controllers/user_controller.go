package controllers

import (
	"go-todo-api/initializers"
	"go-todo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary User List
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
