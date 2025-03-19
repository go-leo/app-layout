#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v protoc-gen-go) ]
then
  echo "--- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest ---"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc-gen-go --version
fi

if [ ! $(command -v protoc-gen-go-grpc) ]
then
  echo "--- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest ---"
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc-gen-go-grpc --version
fi

if [ ! $(command -v protoc-gen-validate-go) ]
then
  echo "--- go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest ---"
	go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest
fi

if [ ! $(command -v protoc-gen-go-leo) ]
then
  echo "--- go install github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo@latest ---"
	go install github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo@latest
	protoc-gen-go-leo --version
fi

echo "--- protoc generate start ---"
protoc \
  --proto_path=. \
  --proto_path=../../third_party \
  --go_out=ui \
  --go_opt=module=qimao/leo/app/user/api/v1 \
  --validate_out=ui \
  --validate_opt=module=qimao/leo/app/user/api/v1,lang=go \
  --go-grpc_out=ui \
  --go-grpc_opt=module=qimao/leo/app/user/api/v1 \
  --go-leo_out=ui \
  --go-leo_opt=module=qimao/leo/app/user/api/v1 \
  api/*/*.proto
echo "--- protoc generate end ---"
