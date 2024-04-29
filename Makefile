
dev:
	@echo "Starting development server..."
	@go run ./terraspect_server/main.go
	@echo "API Running."
	@cd ./terraspect_web && yarn dev
