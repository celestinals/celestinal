protoc --go_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
    --go-tickex_out=. --go-tickex_opt=paths=source_relative \
    example.proto