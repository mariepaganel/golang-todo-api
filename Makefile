# Makefile for Dockerized Echo application

# Variables
DOCKER_COMPOSE = docker compose

# Default target
help: ## Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

# Build and start the Docker containers
up: ## Build and start the Docker containers
	$(DOCKER_COMPOSE) up -d

# Stop the Docker containers
down: ## Stop the Docker containers
	$(DOCKER_COMPOSE) down

# Restart the Docker containers
restart: ## Restart the Docker containers
	$(DOCKER_COMPOSE) restart

# Enter the Docker container shell
shell: ## Enter the Docker container shell
	$(DOCKER_COMPOSE) exec echo sh

# Run a command inside the Docker container
run: ## Run a command inside the Docker container
	$(DOCKER_COMPOSE) exec echo $(CMD)

# Tail logs of the Docker containers
logs: ## Tail logs of the Docker containers
	$(DOCKER_COMPOSE) logs -f

# Remove stopped containers and networks created by `up`
clean: ## Remove stopped containers and networks created by 'up'
	$(DOCKER_COMPOSE) down -v --remove-orphans
