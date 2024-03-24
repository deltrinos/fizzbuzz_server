# Makefile for Fizz-Buzz REST Server

# Build the Fizz-Buzz server
.PHONY: build
build:
	go build -o fizzbuzz_server cmd/main.go

.PHONY: test
# Run unit tests
test:
	go test -v ./...

.PHONY: run
# Run the Fizz-Buzz server
run: build
	./fizzbuzz_server

.PHONY: clean
# Clean up generated files
clean:
	rm -f fizzbuzz_server

.PHONY: format
# Format code
format:
	go fmt ./...

.PHONY: lint
# Run linters
lint:
	golangci-lint run

.PHONY: validate
# Validate code
validate:
	go vet ./...
	golangci-lint run

.PHONY: docker-build
# Docker build the Fizz-Buzz server
docker-build:
	docker build -t fizzbuzz-server .

.PHONY: docker-run
# Docker run the Fizz-Buzz server
docker-run: docker-build
	docker run -p 8000:8000 fizzbuzz-server


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help