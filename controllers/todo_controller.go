package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tofustream/gindo/database"
	"github.com/tofustream/gindo/models"
)

// ToDoリストを取得するエンドポイント
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// 新しいToDoを作成するエンドポイント
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// 指定したIDのToDoを更新するエンドポイント
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// 指定したIDのToDoを削除するエンドポイント
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "ToDo deleted"})
}
