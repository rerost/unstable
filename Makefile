.PHONY: protoc
protoc: 
	protoc --proto_path=./common/internal/proto --go_out=./common/internal/proto ./common/internal/proto/config.proto
