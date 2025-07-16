package models

import (
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status" validate:"oneof:In Progress Completed"`
   }