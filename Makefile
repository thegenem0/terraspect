API_DIR=./terraspect_server
WEB_DIR=./terraspect_web

.PHONY: dev api web build-api build-web clean

dev: api web
	@echo "Both services are running..."

api:
	@echo "Running Go backend..."
	cd $(API_DIR) && go run . &

web:
	@echo "Running Vite frontend..."
	cd $(WEB_DIR) && yarn dev &

build-api:
	@echo "Building Go backend..."
	cd $(API_DIR) && go build -o app

build-web:
	@echo "Building Vite frontend..."
	cd $(WEB_DIR) && npm run build

clean:
	@echo "Cleaning up..."
	cd $(API_DIR) && go clean
	cd $(WEB_DIR) && rm -rf dist

stop:
	@echo "Stopping all services..."
	-@pkill -P $$!