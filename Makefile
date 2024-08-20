run:
	@echo "go mod tidy..."
	@go mod tidy
	@echo "go mod vendor..."
	@go mod vendor
	@echo "Running my application..."
	@air
