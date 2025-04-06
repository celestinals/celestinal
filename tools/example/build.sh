protoc --go_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
    --go-celestinal_out=. --go-celestinal_opt=paths=source_relative \
    example.proto