.PHONY: build
build:
	@echo "Building now."
	@go build ./src/main/webapi/main.go ./src/main/webapi/person.go

.PHONY: run
run: build
	@echo "run"
	@go run ./src/main/webapi/main.go ./src/main/webapi/person.go
