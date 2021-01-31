package handler

import (
	"context"

	greeter "github.com/kzmake/dapr-kit/example/helloworld/api/greeter/v1"
)

// Greeter ...
type Greeter struct {
	greeter.UnimplementedGreeterServiceServer
}

// interfaces
var _ greeter.GreeterServiceServer = (*Greeter)(nil)
var _ greeter.GreeterServiceHandler = (*Greeter)(nil)

// Hello ...
func (c *Greeter) Hello(_ context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{Msg: "hello, " + req.Name}, nil
}
