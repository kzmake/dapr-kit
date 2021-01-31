package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	bob "github.com/kzmake/dapr-kit/example/request-replay/api/bob/v1"

	"github.com/kzmake/dapr-kit/example/request-replay/bob/handler"
)

func main() {
	s, err := daprd.NewService(":4002")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	b := &handler.Bob{}
	if err := bob.RegisterBobServiceInvocationHandlers(s, b); err != nil {
		log.Fatalf("failed to register invocation handlers: %v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
