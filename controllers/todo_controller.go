package controllers

import (
	"go-todo-api/dto"
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

// @Summary Create ToDo
// @Description Create a ToDo
// @Tags ToDo
// @Accept       json
// @Produce      json
// @Param	data	body	dto.ToDoAttrs true "ToDo Attributes"
// @Success 200 {object} models.ToDo
// @Failure      401  {object} dto.ErrorResponse
// @Failure      400  {object} dto.ErrorResponse
// @Router /todo [post]
// @Security Bearer
func ToDoCreate(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	body := dto.ToDoAttrs{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.ToDo{
		Task:      body.Task,
		Completed: body.Completed,
		UserId:    &userId,
	}

	initializers.DB.Save(&todo)

	c.JSON(http.StatusCreated, todo)
}

// @Summary ToDo by ID
// @Description Get a ToDo by ID
// @Tags ToDo
// @Accept       json
// @Produce      json
// @Param ToDoID path int true "ToDo ID"
// @Success 200 {object} models.ToDo
// @Failure      401  {object} dto.ErrorResponse
// @Router /todo/{ToDoID} [get]
// @Security Bearer
func ToDoById(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	todo := models.ToDo{}

	initializers.DB.First(&todo, c.Param("id"))

	c.JSON(http.StatusOK, todo)
}

// @Summary Update ToDo
// @Description Update a ToDo
// @Tags ToDo
// @Accept       json
// @Produce      json
// @Param ToDoID path int true "ToDo ID"
// @Param	data	body	dto.ToDoAttrs true "ToDo Attributes"
// @Success 200 {object} models.ToDo
// @Failure      401  {object} dto.ErrorResponse
// @Failure      400  {object} dto.ErrorResponse
// @Router /todo/{ToDoID} [patch]
// @Security Bearer
func ToDoUpdate(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	body := dto.ToDoAttrs{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.ToDo{}

	initializers.DB.First(&todo, c.Param("id"))
	todo.Task = body.Task
	todo.Completed = body.Completed

	initializers.DB.Save(&todo)

	c.JSON(http.StatusOK, todo)
}

// @Summary Delete ToDo
// @Description Delete a ToDo
// @Tags ToDo
// @Accept       json
// @Produce      json
// @Param ToDoID path int true "ToDo ID"
// @Success 204
// @Failure      401  {object} dto.ErrorResponse
// @Router /todo/{ToDoID} [delete]
// @Security Bearer
func ToDoDelete(c *gin.Context) {
	userId := utilities.UserIdFromHeader(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	todo := models.ToDo{}

	initializers.DB.First(&todo, c.Param("id"))
	todo.IsDeleted = true

	initializers.DB.Save(&todo)

	c.Writer.WriteHeader(http.StatusNoContent)
}
