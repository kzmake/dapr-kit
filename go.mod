module github.com/kzmake/dapr-kit

go 1.15

replace github.com/dapr/components-contrib => github.com/dapr/components-contrib v1.0.0

require (
	github.com/dapr/components-contrib v1.0.0
	github.com/dapr/go-sdk v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0
	github.com/json-iterator/go v1.1.10
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)
