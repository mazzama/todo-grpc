# TO DO LIST APP 

## Description
### Database Schema
DDL
```sql
CREATE TABLE IF NOT EXISTS ITEMS (
	id bigserial PRIMARY KEY,
	name varchar(50) not null,
	description varchar(255),
	notes varchar(100),
	status varchar(20) not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP not null
);
```
### Service List
```protobuf
service TodoService {
  rpc CreateItem(CreateItemRequest) returns (CreateItemResponse);
  rpc UpdateItem(UpdateItemRequest) returns (Item);
  rpc DeleteItem(DeleteItemRequest) returns (DeleteItemResponse);
  rpc FindItemByID(FindItemRequest) returns (Item);
  rpc ViewItemList(ViewItemListRequest) returns (ViewItemListResponse);
}
```

## Library
- Viper
- go-grpc
- GORM
- bufconn
- golang-migrate V4
- github.com/stretchr/testify
- Mockery

## Software Requirements
- [Golang 1.19](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [BloomRPC](https://github.com/bloomrpc/bloomrpc/releases/tag/1.5.3)

Optional installation:
- To generate grpc client and server:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
- To generate mock file
```
go install github.com/vektra/mockery@v2.28.2
mockery --dir internal/ --all --output --output=internal/mocks
```

## How to run
- Run unit test
```
make test.unit
```
- Run integration test (need to run make docker.start)
```
make test.integration
```
- Start the application for local test
```
make docker.start
```
Open BloomRPC and import protobuf file from this folder /pkg/proto. \
Adjust client port from config.yaml file
- Stop the application
```
make docker.stop
```