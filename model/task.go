package model

import "time"

type Task struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"             binding:"required,max=100"`
	Done      bool      `json:"done" db:"done"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
