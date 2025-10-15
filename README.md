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
