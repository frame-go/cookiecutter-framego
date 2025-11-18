module {{ cookiecutter.go_module }}

go {{ cookiecutter.go_version }}

require (
)

tool (
	github.com/envoyproxy/protoc-gen-validate
	github.com/frame-go/protoc-gen-framego
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	github.com/swaggo/swag/cmd/swag
	github.com/swaggo/swag/gen
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	google.golang.org/protobuf/cmd/protoc-gen-go
)
