BINARY_NAME=go-todo-api
DOCKER_IMAGE=go-todo-api:latest

.PHONY: build
build:
	@echo "Building Go binary..."
	go build -o $(BINARY_NAME) ./cmd/server

.PHONY: run
run:
	@echo "Running Go application..."
	go run ./cmd/server

.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./tests/...

.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

.PHONY: k8s-deploy
k8s-deploy: docker-build
	@echo "Deploying to Kubernetes..."
	kubectl apply -f k8s/deployment.yaml

.PHONY: k8s-delete
k8s-delete:
	@echo "Deleting from Kubernetes..."
	kubectl delete -f k8s/deployment.yaml

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build         - Build Go binary"
	@echo "  run           - Run Go application"
	@echo "  test          - Run tests"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  k8s-deploy    - Deploy to Kubernetes"
	@echo "  k8s-delete    - Delete from Kubernetes"
	@echo "  clean         - Clean up"
