 docker-compose up:
	docker-compose -f docker-compose-developer.yml up --build -d
 docker-compose down:
	docker-compose -f docker-compose-developer.yml down

generate-proto:
	protoc --go_out=./proto/gen/pb_detects --go-grpc_out=./proto/gen/pb_detects ./proto/detect/detect.proto