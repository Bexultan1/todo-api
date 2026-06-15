package handler

import (
	"time"
	"todo-api/db"
	"todo-api/model"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []model.Task

	err := db.DB.Select(&tasks, "SELECT * FROM tasks")

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	var task model.Task
	err := db.DB.Get(&task, "SELECT * FROM tasks WHERE ID =$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, task)
}

func CreateTask(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task.CreatedAt = time.Now()

	err := db.DB.QueryRow("INSERT INTO tasks (title, done, created_at) VALUES ($1,$2,$3) RETURNING id", task.Title, task.Done, task.CreatedAt).Scan(&task.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task model.Task

	if err := db.DB.Get(&task, "SELECT * FROM tasks WHERE ID =$1", id); err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec(
		"UPDATE tasks SET title = $1, done = $2 WHERE ID = $3", task.Title, task.Done, task.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task

	if err := db.DB.Get(&task, "SELECT * FROM tasks WHERE ID =$1", id); err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	_, err := db.DB.Exec("DELETE FROM tasks WHERE ID =$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "Task deleted"})
}
