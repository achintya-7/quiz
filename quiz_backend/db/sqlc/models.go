// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"
)

type Question struct {
	ID        int64     `json:"id"`
	QuizID    int64     `json:"quiz_id"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
	Answer    string    `json:"answer"`
}

type Quiz struct {
	ID          int64     `json:"id"`
	CreatedBy   string    `json:"created_by"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}