package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tofustream/gindo/controllers"
	"github.com/tofustream/gindo/database"
)

func main() {
	database.Connect()

	// Ginエンジンのインスタンスを生成
	router := gin.Default()

	// CORSの設定を追加
	router.Use(cors.Default())

	// エンドポイントの設定
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	// サーバー起動
	router.Run(":8080")
}
