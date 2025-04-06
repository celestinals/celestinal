module github.com/celestinals/celestinal/staging/mod/apigateway

go 1.24.2

replace (
	github.com/celestinals/celestinal => ../../../
	github.com/celestinals/celestinal/api => ../../../api
)

require (
	github.com/celestinals/celestinal v0.0.0-00010101000000-000000000000
	github.com/celestinals/celestinal/api v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	google.golang.org/grpc v1.71.1
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.5-20250219170025-d39267d9df8f.1 // indirect
	cel.dev/expr v0.20.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/bufbuild/protovalidate-go v0.9.2 // indirect
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // indirect
	github.com/google/cel-go v0.23.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	go.uber.org/dig v1.18.0 // indirect
	go.uber.org/fx v1.23.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/exp v0.0.0-20250218142911-aa4b98e5adaa // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
