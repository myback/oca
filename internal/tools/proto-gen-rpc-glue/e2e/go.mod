module github.com/myback/oca/internal/tools/proto-gen-rpc-glue/e2e

go 1.21

toolchain go1.23.2

replace github.com/myback/oca => ./consul

require github.com/myback/oca v0.0.0-00010101000000-000000000000

require google.golang.org/protobuf v1.35.2 // indirect
