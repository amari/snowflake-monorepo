module github.com/amari/snowflake-monorepo/examples/go

go 1.24.4

require (
	github.com/amari/snowflake-monorepo v0.0.0-20251020030136-87bc2c73a1dd
	google.golang.org/grpc v1.76.0
)

require (
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251007200510-49b9836ed3ff // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/amari/snowflake-monorepo v0.0.0-20251020030136-87bc2c73a1dd => ../../
