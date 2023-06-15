build:
	docker build -t trevatk/healthcheck:latest .

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run
