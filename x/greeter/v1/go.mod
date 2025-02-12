module github.com/tickexvn/tickex/x/greeter/v1

go 1.23.3

replace (
	github.com/tickexvn/tickex => ./../../../
	github.com/tickexvn/tickex/api => ./../../../api
)

require (
	github.com/jackc/pgx/v5 v5.7.2
	github.com/tickexvn/tickex v0.0.0-00010101000000-000000000000
	github.com/tickexvn/tickex/api v0.0.0-00010101000000-000000000000
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.2-20241127180247-a33202765966.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	go.uber.org/dig v1.18.0 // indirect
	go.uber.org/fx v1.23.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
)
