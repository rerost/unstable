.PHONY: protoc
protoc: 
	protoc --proto_path=./initialize/internal/proto --go_out=./initialize/internal/proto ./initialize/internal/proto/config.proto
