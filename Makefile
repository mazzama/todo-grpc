.PHONY: compile
compile: ## Compile the proto file.
	protoc --proto_path=pkg/proto pkg/proto/*.proto --go_out=. --go-grpc_out=.

.PHONY: run
run: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server cmd/main.go
	bin/server

.PHONY: test.unit
test.unit: ## run unit test
	go test ./...

.PHONY: test.integration
test.integration: docker.app.stop
	go test -tags=integration ./it -v -count=1

.PHONY: docker.start
docker.start:
		docker-compose up -d

.PHONY: docker.stop
docker.stop:
		docker-compose down -v

.PHONY: docker.app.stop
docker.app.stop:
		docker stop todo-app

