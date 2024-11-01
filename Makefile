include local.env

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.22.1


install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.2

get-deps:
		go get github.com/golang/protobuf/ptypes/empty
		go get github.com/golang/protobuf/ptypes/timestamp
	 	go get google.golang.org/protobuf/reflect/protoreflect
		go get google.golang.org/protobuf/runtime/protoimpl
		go get google.golang.org/grpc
		go get google.golang.org/grpc/codes
		go get google.golang.org/grpc/status
		 go get github.com/brianvoe/gofakeit
		 go get github.com/jackc/pgx/v5


generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

migrations-up :
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(MIGRATION_DSN_L) ./bin/goose -dir migrations up -v
migrations-down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(MIGRATION_DSN_L) ./bin/goose -dir migrations down -v

docker-postgres:
	 docker run  --rm -d --name DB -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db-chat -p 54341:5432 postgres:14-alpine3.20
	 docker build --tag chat_migrator_image -f  migration_local.Dockerfile .