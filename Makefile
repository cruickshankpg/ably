protos:
	protoc --go_out=plugins=grpc:. protos/stateful/*.proto
	protoc --go_out=plugins=grpc:. protos/stateless/*.proto

build: build-stateless build-stateful

build-stateless:
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	 go build \
		-a \
		-o stateless-server \
		./cmd/statelessserver

build-stateful:
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	 go build \
		-a \
		-o stateful-server \
		./cmd/statefulserver