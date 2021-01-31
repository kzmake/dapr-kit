package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	alice "github.com/kzmake/dapr-kit/example/request-replay/api/alice/v1"

	"github.com/kzmake/dapr-kit/example/request-replay/alice/handler"
)

func main() {
	s, err := daprd.NewService(":4001")
	if err != nil {
		log.Fatalf("failed to create the server: %v", err)
	}

	a, err := handler.NewAlice("bob", "localhost:3500")
	if err != nil {
		log.Fatalf("failed to create the handler: %v", err)
	}
	if err := alice.RegisterAliceServiceInvocationHandlers(s, a); err != nil {
		log.Fatalf("failed to register invocation handlers: %v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
