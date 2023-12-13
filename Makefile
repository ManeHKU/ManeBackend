gen_proto:
	echo "Building Proto"
	protoc --proto_path=./../protobuf --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative service.proto