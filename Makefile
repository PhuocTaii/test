server:
	@echo "Starting server..."
	@go run cmd/server/main.go

lint:
	@echo "golangci lint checking..."
	@golangci-lint run

tidy:
	@sudo rm -rf vendor
	@go mod tidy
	@go mod vendor

compose_up:
	@docker compose up -d

compose_down:
	@docker compose down
	@docker rmi -f ia_03:latest 

gen_doc:
	@swag init --parseDependency -g cmd/server/main.go
