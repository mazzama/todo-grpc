.PHONY: compile
compile: ## Compile the proto file.
	protoc --proto_path=pkg/proto pkg/proto/*.proto --go_out=. --go-grpc_out=.

.PHONY: run
run: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server cmd/main.go
	bin/server

test.unit: ## run unit test
	go test ./...

test.integration:  ## run integration test
	go test -tags=integration ./it -v -count=1