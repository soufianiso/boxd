include .env
run:
	@go mod tidy
	@echo "go mod vendor..."
	@go mod vendor
	@echo "Running my application..."
	@air

migrate:
	@echo ${postgres}
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(postgres) goose -dir=db/migrations up

down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(postgres) goose -dir=db/migrations down



