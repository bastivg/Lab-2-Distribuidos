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
	sudo docker build . -t australia:1.21.1

docker-continentes:
	sudo docker run --name australia australia:1.21.1