.PHONY: help

dev: compose-up
stop: compose-down
deploy: build-web deploy-stack

compose-down:
	@docker compose down

compose-up:
	@docker compose up

api:
	@docker compose up postgres api

web:
	@docker compose up web

build-web:
	@cd terraspect_web && yarn build

deploy-stack:
	@cd infra && cdk deploy

swagger:
	@cd terraspect_server && swag init

help:
	@echo "dev: Run the development environment"
	@echo "stop: Stop the development environment"
	@echo "deploy: Deploy the stack"
	@echo "api: Run the api service"
	@echo "web: Run the web service"
	@echo "build-web: Build the web service"
	@echo "deploy-stack: Deploy the stack"
	@echo "swagger: Generate swagger documentation"
	@echo "help: Show this help message"