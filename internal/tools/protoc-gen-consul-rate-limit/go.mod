module github.com/myback/oca/internal/tools/protoc-gen-consul-rate-limit

go 1.21

toolchain go1.23.2

replace github.com/myback/oca/proto-public => ../../../proto-public

require (
	github.com/myback/oca/proto-public v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.35.2
)

require github.com/google/go-cmp v0.5.9 // indirect
