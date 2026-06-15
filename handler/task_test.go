package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"todo-api/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	godotenv.Load("../.env")
	db.Connect()

	r := gin.Default()
	r.GET("/tasks", GetTasks)

	req, _ := http.NewRequest("GET", "/tasks", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateTask(t *testing.T) {
	godotenv.Load("../.env")
	db.Connect()

	r := gin.Default()
	r.POST("/tasks", CreateTask)

	body := strings.NewReader(`{"title":"тест","done":false}`)
	req, _ := http.NewRequest("POST", "/tasks", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

}
