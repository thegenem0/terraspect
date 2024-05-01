
dev: compose-up

stop: compose-down

compose-down:
	@docker compose down

compose-up:
	@docker compose up

api:
	@docker compose up postgres api

web:
	@docker compose up web

swagger:
	@cd terraspect_server && swag init