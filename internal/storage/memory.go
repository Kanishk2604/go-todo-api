package storage

import (
	"errors"
	"go-todo-api/internal/models"
	"sync"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type TodoStorage interface {
	Create(todo *models.Todo) error
	GetByID(id string) (*models.Todo, error)
	GetAll() ([]*models.Todo, error)
	Update(id string, req models.UpdateTodoRequest) (*models.Todo, error)
	Delete(id string) error
}

type InMemoryStorage struct {
	todos map[string]*models.Todo
	mutex sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		todos: make(map[string]*models.Todo),
	}
}

func (s *InMemoryStorage) Create(todo *models.Todo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.todos[todo.ID] = todo
	return nil
}

func (s *InMemoryStorage) GetByID(id string) (*models.Todo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	todo, exists := s.todos[id]
	if !exists {
		return nil, ErrTodoNotFound
	}
	return todo, nil
}

func (s *InMemoryStorage) GetAll() ([]*models.Todo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	todos := make([]*models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *InMemoryStorage) Update(id string, req models.UpdateTodoRequest) (*models.Todo, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	todo, exists := s.todos[id]
	if !exists {
		return nil, ErrTodoNotFound
	}
	
	todo.Update(req)
	return todo, nil
}

func (s *InMemoryStorage) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	if _, exists := s.todos[id]; !exists {
		return ErrTodoNotFound
	}
	
	delete(s.todos, id)
	return nil
}
