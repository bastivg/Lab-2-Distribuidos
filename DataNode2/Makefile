generate_grpc_code:
	cd proto && \
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	service.proto && \
	cd ..

build-docker:
	sudo docker build . -t datanodetwo:1.21.1

docker-datanode:
	sudo docker run --name datanodetwo datanodetwo:1.21.1