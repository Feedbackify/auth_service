SWAGGER_UI_VERSION:=v4.15.5
BUF_VERSION:=v1.17.0
BASE_IMAGE_NAME := feedbackify
SERVICE_NAME    := auth_service
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)


install:
	go get \
	    github.com/bufbuild/buf/cmd/buf \
 		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
 		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
 		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
 		google.golang.org/protobuf/cmd/protoc-gen-go

run:
	go run cmd/main.go
dev:
	go run main.go


generate: generate/proto
generate/proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

lint:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) lint
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) breaking --against 'https://github.com/johanbrandhorst/grpc-gateway-boilerplate.git#branch=main'

docker:
	docker build \
    		-t $(SERVICE_IMAGE) \
    		--build-arg BUILD_REF=$(VERSION) \
    		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
    		.