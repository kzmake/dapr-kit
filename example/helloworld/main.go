package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	greeter "github.com/kzmake/dapr-kit/example/helloworld/api/greeter/v1"

	"github.com/kzmake/dapr-kit/example/helloworld/handler"
)

func main() {
	// create a Dapr service server
	s, err := daprd.NewService(":50001")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	g := &handler.Greeter{}
	if err := greeter.RegisterGreeterServiceInvocationHandlers(s, g); err != nil {
		log.Fatalf("failed to register invocation handlers: %v", err)
	}

	// start the server
	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
