package main

import (
	"todo-api/db"
	"todo-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.GET("/tasks", handler.GetTasks)

	r.GET("/tasks/:id", handler.GetTask)

	r.POST("/tasks", handler.CreateTask)

	r.PUT("/tasks/:id", handler.UpdateTask)

	r.DELETE("/tasks/:id", handler.DeleteTask)

	r.Run(":8080")
}
