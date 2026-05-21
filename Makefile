.PHONY: help dev build test lint clean docker-up docker-down

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Start both frontend and backend in dev mode
	@echo "Starting backend..."
	@cd backend && make watch &
	@echo "Starting frontend..."
	@cd frontend && npm run dev

build: ## Build both frontend and backend
	@echo "Building frontend..."
	@cd frontend && npm run build
	@echo "Building backend..."
	@cd backend && make build

test: ## Run all tests
	@echo "Running frontend tests..."
	@cd frontend && npx vitest run
	@echo "Running backend tests..."
	@cd backend && go test ./... -v

lint: ## Lint both frontend and backend
	@echo "Linting frontend..."
	@cd frontend && npm run lint
	@echo "Linting backend..."
	@cd backend && go vet ./...

clean: ## Clean build artifacts
	@cd frontend && rm -rf dist
	@cd backend && make clean

docker-up: ## Start all services with docker compose
	@if docker compose up -d 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up -d; \
	fi

docker-down: ## Stop all services
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi
