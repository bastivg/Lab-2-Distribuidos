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
	sudo docker build . -t oms:1.21.1

docker-OMS:
	sudo docker run --name oms -p 50053:50053 oms:1.21.1
