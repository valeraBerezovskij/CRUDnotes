package models

import "time"

type Note struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
