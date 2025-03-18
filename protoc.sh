protoc \
		--proto_path=. \
		--proto_path=../../third_party \
		--go_out=ui \
		--go_opt=module=github.com/go-leo/app-layout/api/v1 \
		--validate_out=ui \
		--validate_opt=module=github.com/go-leo/app-layout/api/v1,lang=go \
		--go-grpc_out=ui \
    --go-grpc_opt=module=github.com/go-leo/app-layout/api/v1 \
		--go-leo_out=ui \
		--go-leo_opt=module=github.com/go-leo/app-layout/api/v1 \
		api/*/*.proto