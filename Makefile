protos:
	protoc --go_out=plugins=grpc:. protos/stateful/*.proto
	protoc --go_out=plugins=grpc:. protos/stateless/*.proto