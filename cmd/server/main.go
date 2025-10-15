package main

import (
	"go-todo-api/internal/handlers"
	"go-todo-api/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	logger := log.New(os.Stdout, "[TODO-API] ", log.LstdFlags|log.Lshortfile)
	todoStorage := storage.NewInMemoryStorage()
	todoHandler := handlers.NewTodoHandler(todoStorage, logger)
	
	r := chi.NewRouter()
	
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
	
	r.Route("/api/v1/todos", func(r chi.Router) {
		r.Post("/", todoHandler.CreateTodo)
		r.Get("/", todoHandler.GetTodos)
		r.Get("/{id}", todoHandler.GetTodo)
		r.Put("/{id}", todoHandler.UpdateTodo)
		r.Delete("/{id}", todoHandler.DeleteTodo)
	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	logger.Printf("Starting server on port %s", port)
	logger.Printf("Health check: http://localhost:%s/health", port)
	logger.Printf("API: http://localhost:%s/api/v1/todos", port)
	
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Fatal("Server failed to start:", err)
	}
}
