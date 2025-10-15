package models

import (
	"time"
	"github.com/google/uuid"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
}

func NewTodo(title, description string) *Todo {
	now := time.Now()
	return &Todo{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Todo) Update(req UpdateTodoRequest) {
	if req.Title != nil {
		t.Title = *req.Title
	}
	if req.Description != nil {
		t.Description = *req.Description
	}
	if req.Completed != nil {
		t.Completed = *req.Completed
	}
	t.UpdatedAt = time.Now()
}
