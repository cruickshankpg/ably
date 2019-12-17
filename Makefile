protos: stateless-proto stateful-proto

stateful-proto:
	protoc --go_out=plugins=grpc:. protos/stateful/*.proto

stateless-proto:
	protoc --go_out=plugins=grpc:. protos/stateless/*.proto

build: build-stateless build-stateful

build-stateless:
	 go build \
		-a \
		-o stateless-server \
		./cmd/statelessserver

build-stateful:
	 go build \
		-a \
		-o stateful-server \
		./cmd/statefulserver

