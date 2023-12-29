gen-proto:
	cd ../protobuf; git checkout main && git pull
	protoc --proto_path=./../protobuf --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative main_service.proto
	protoc --proto_path=./../protobuf --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative health_check.proto
	protoc --proto_path=./../protobuf --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative init_service.proto
run:
	echo "Running gRPC"
	go run ./cmd/main.go