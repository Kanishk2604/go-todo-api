package handlers

import (
	"encoding/json"
	"errors"
	"go-todo-api/internal/models"
	"go-todo-api/internal/storage"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TodoHandler struct {
	storage storage.TodoStorage
	logger  *log.Logger
}

func NewTodoHandler(storage storage.TodoStorage, logger *log.Logger) *TodoHandler {
	return &TodoHandler{
		storage: storage,
		logger:  logger,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *TodoHandler) respondWithError(w http.ResponseWriter, r *http.Request, code int, message string) {
	h.logger.Printf("Error: %s - %s %s", message, r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, r, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if req.Title == "" {
		h.respondWithError(w, r, http.StatusBadRequest, "Title is required")
		return
	}

	todo := models.NewTodo(req.Title, req.Description)
	if err := h.storage.Create(todo); err != nil {
		h.respondWithError(w, r, http.StatusInternalServerError, "Failed to create todo")
		return
	}

	h.logger.Printf("Created todo: %s", todo.ID)
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, todo)
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.storage.GetAll()
	if err != nil {
		h.respondWithError(w, r, http.StatusInternalServerError, "Failed to retrieve todos")
		return
	}

	h.logger.Printf("Retrieved %d todos", len(todos))
	render.JSON(w, r, todos)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondWithError(w, r, http.StatusBadRequest, "Todo ID is required")
		return
	}

	todo, err := h.storage.GetByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrTodoNotFound) {
			h.respondWithError(w, r, http.StatusNotFound, "Todo not found")
			return
		}
		h.respondWithError(w, r, http.StatusInternalServerError, "Failed to retrieve todo")
		return
	}

	h.logger.Printf("Retrieved todo: %s", id)
	render.JSON(w, r, todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondWithError(w, r, http.StatusBadRequest, "Todo ID is required")
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, r, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	todo, err := h.storage.Update(id, req)
	if err != nil {
		if errors.Is(err, storage.ErrTodoNotFound) {
			h.respondWithError(w, r, http.StatusNotFound, "Todo not found")
			return
		}
		h.respondWithError(w, r, http.StatusInternalServerError, "Failed to update todo")
		return
	}

	h.logger.Printf("Updated todo: %s", id)
	render.JSON(w, r, todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondWithError(w, r, http.StatusBadRequest, "Todo ID is required")
		return
	}

	err := h.storage.Delete(id)
	if err != nil {
		if errors.Is(err, storage.ErrTodoNotFound) {
			h.respondWithError(w, r, http.StatusNotFound, "Todo not found")
			return
		}
		h.respondWithError(w, r, http.StatusInternalServerError, "Failed to delete todo")
		return
	}

	h.logger.Printf("Deleted todo: %s", id)
	w.WriteHeader(http.StatusNoContent)
}
