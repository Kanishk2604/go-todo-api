package tests

import (
	"encoding/json"
	"go-todo-api/internal/handlers"
	"go-todo-api/internal/models"
	"go-todo-api/internal/storage"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func setupTestHandler() *handlers.TodoHandler {
	logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
	storage := storage.NewInMemoryStorage()
	return handlers.NewTodoHandler(storage, logger)
}

func setupRouter(handler *handlers.TodoHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/todos", func(r chi.Router) {
		r.Post("/", handler.CreateTodo)
		r.Get("/", handler.GetTodos)
		r.Get("/{id}", handler.GetTodo)
		r.Put("/{id}", handler.UpdateTodo)
		r.Delete("/{id}", handler.DeleteTodo)
	})
	return r
}

func TestCreateTodo(t *testing.T) {
	handler := setupTestHandler()
	router := setupRouter(handler)

	tests := []struct {
		name           string
		payload        string
		expectedStatus int
		expectError    bool
	}{
		{
			name:           "Valid todo creation",
			payload:        `{"title":"Test Todo","description":"Test Description"}`,
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name:           "Missing title",
			payload:        `{"description":"Test Description"}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "Empty title",
			payload:        `{"title":"","description":"Test Description"}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/api/v1/todos", strings.NewReader(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if !tt.expectError {
				var todo models.Todo
				if err := json.Unmarshal(rr.Body.Bytes(), &todo); err != nil {
					t.Errorf("Failed to parse todo response: %v", err)
				}
				if todo.ID == "" {
					t.Error("Expected todo ID to be set")
				}
			}
		})
	}
}
