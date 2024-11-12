package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tofustream/gindo/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// ToDoのエンドポイント
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	return r
}
