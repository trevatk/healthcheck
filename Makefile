build:
	docker build -t trevatk/healthcheck:v0.0.1 .

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run
