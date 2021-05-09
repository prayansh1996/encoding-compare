PROTOC_DOCKER_IMAGE = public.ecr.aws/zomato/zomato/protoc-gen:v2.0

compile_protos:
	@echo "Compiling protos:\n"
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) ${PROTOC_DOCKER_IMAGE} --go_out=plugins=grpc:. ./cache/base.proto

run: compile_protos
	@echo "\n\nSerialized cache size comparision:\n"
	@go run main.go

	@echo "\n\nRunning performance benchmarks:"
	@echo "Comparing MsgPack and Snappy with Protobuf encoding\n"
	@go test -bench=. ./cache