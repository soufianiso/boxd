include .env
# ruw q:
# 	@go mod tidy
# 	@echo "go mod vendor..."
# 	@go mod vendor
# 	@echo "Running my application..."
# 	@air

migrate:
	@echo ${postgres}
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(postgres) goose -dir=db/migrations up

down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(postgres) goose -dir=db/migrations down



watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi


