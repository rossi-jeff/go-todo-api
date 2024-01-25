package controllers

import (
	"go-todo-api/initializers"
	"go-todo-api/models"
	"go-todo-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary ToDo List
// @Description List of ToDos for logged in user
// @Tags ToDo
// @Accept       json
// @Produce      json
// @Success 200 {array} models.ToDo
// @Failure      401  {object} dto.ErrorResponse
// @Router /todo [get]
// @Security Bearer
func ToDoList(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	todos := []models.ToDo{}

	initializers.DB.Model(&models.ToDo{}).Where("user_id = ? and is_deleted = ?", userId, false).Find(&todos)

	c.JSON(http.StatusOK, todos)
}
