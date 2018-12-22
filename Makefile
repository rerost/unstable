.PHONY: protoc
protoc: 
	protoc --proto_path=./init/proto --go_out=./init/proto ./init/proto/config.proto
