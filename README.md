# Go Todo REST API

A professional Todo REST API built with Go, featuring clean architecture, comprehensive testing, and container deployment capabilities.

## Features

✅ **Clean Architecture** - Separation of handlers, models, and storage layers  
✅ **CRUD Operations** - Complete Create, Read, Update, Delete functionality  
✅ **Structured Logging** - Request tracing and error handling  
✅ **Table-Driven Tests** - Comprehensive test coverage  
✅ **Docker Ready** - Multi-stage containerization  
✅ **Kubernetes Ready** - Production deployment manifests  
✅ **Thread-Safe** - Concurrent request handling with mutex protection  

## Tech Stack

- **Language**: Go 1.21
- **Router**: Chi v5 (lightweight, fast HTTP router)
- **Storage**: In-memory with thread-safe operations
- **Testing**: Go standard testing with table-driven tests
- **Containerization**: Docker multi-stage builds
- **Orchestration**: Kubernetes with health checks

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| POST | `/api/v1/todos` | Create a new todo |
| GET | `/api/v1/todos` | Get all todos |
| GET | `/api/v1/todos/{id}` | Get specific todo by ID |
| PUT | `/api/v1/todos/{id}` | Update todo by ID |
| DELETE | `/api/v1/todos/{id}` | Delete todo by ID |

## Quick Start

### Prerequisites
- Go 1.21 or later
- Docker (optional)
- Kubernetes cluster (optional)

### Run Locally
1. **Clone the repository**
```bash
git clone https://github.com/Kanishk2604/go-todo-api.git
cd go-todo-api
```
2. **Install dependencies**
```bash
go mod tidy
```
3. **Run the server**
   ```bash
   go run ./cmd/server
   ```
4. **Health check**
   ```bash
   curl http://localhost:8080/health
   ```
5. **Create a todo**
   ```bash
   curl -X POST http://localhost:8080/api/v1/todos
   -H "Content-Type:application/json"
   -d'{"title":"Learn Go","description":"Master Go development"}'
   ```
6. **Get all todos
   ```bash
   curl http://localhost:8080/api/v1/todos
   ```
   
## Testing

Run the comprehensive test suite:
```bash
go test -v ./tests/...
```

Features table-driven tests covering:
- Valid input validation
- Error handling scenarios
- HTTP status code verification
- JSON response parsing

## Docker Deployment

1. **Build Docker image**
   ```bash
   docker build -t go-todo-api:latest .
   ```
2. **Run container**
   ```bash
   docker run -p 8080:8080 go-todo-api:latest
   ```
   
## Kubernetes Deployment

1. **Deploy to Kubernetes**
   ```bash
   kubectl apply -f k8s/deployment.yaml
   ```
2. **Check status**
   ```bash
   kubectl get pods -l app=go-todo-api
   kubectl get services -l app=go-todo-api
   ```
   
