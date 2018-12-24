.PHONY: protoc
protoc: 
	protoc --proto_path=./common/internal/proto --go_out=./common/internal/proto ./common/internal/proto/config.proto

.PHONY: setup
setup:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@dep ensure

.PHONY: test
test: setup
	@go test -v ./...
