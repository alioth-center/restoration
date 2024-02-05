BASE_PATH=$(shell pwd)
PROTO_PATH=$(BASE_PATH)/proto
PROTO_CMD=protoc
GO_OUT_ARG=.
GO_GRPC_OUT_ARG=.
PROTO_GEN_FILES_ARG=./*.proto

DEV_DATA_PATH=$(BASE_PATH)/data
DEV_IMAGE_NAME=github.com/alioth-center/restoration:local-dev
DEV_CONTAINER_NAME=restoration-dev-server

clean_proto:
	rm -rf $(PROTO_PATH)/**/*.pb.go

generate_proto:
	cd $(PROTO_PATH) && $(PROTO_CMD) --go_out=$(GO_OUT_ARG) --go-grpc_out=$(GO_GRPC_OUT_ARG) $(PROTO_GEN_FILES_ARG)

clean_dev_data:
	rm -rf $(DEV_DATA_PATH) && mkdir -p $(DEV_DATA_PATH)

generate: clean_proto generate_proto

docker_build:
	cd $(BASE_PATH) && docker build . -t $(DEV_IMAGE_NAME)

docker_run:
	docker run -p 28081:28081 -v $(DEV_DATA_PATH):/app/data --name $(DEV_CONTAINER_NAME) -d --rm $(DEV_IMAGE_NAME)

docker_stop:
	docker ps -a | grep "$(DEV_CONTAINER_NAME)" && docker rm -f "$(DEV_CONTAINER_NAME)" || echo "No container to stop"

docker_clean:
	docker images | grep "$(DEV_IMAGE_NAME)" && docker rmi -f "$(DEV_IMAGE_NAME)" || echo "No image to clean"

docker_start: docker_stop docker_clean docker_build docker_run