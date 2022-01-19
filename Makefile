.PHONY: setup
setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: protoc
protoc:
	protoc \
		-Iproto \
		--go_out=plugins=grpc:. \
		proto/*.proto

.PHONY: run
run:
	cd grpc && fresh

.PHONY: server
server:
	go run server/grpc/server.go

.PHONY: client
client:
	go run cmd/main.go
