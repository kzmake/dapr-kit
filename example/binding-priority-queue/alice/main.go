package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	alice "github.com/kzmake/dapr-kit/example/binding-priority-queue/api/alice/v1"

	"github.com/kzmake/dapr-kit/example/binding-priority-queue/alice/handler"
)

const (
	appAddress  = ":4000"
	daprAddress = ":3500"
)

func main() {
	s, err := daprd.NewService(appAddress)
	if err != nil {
		log.Fatalf("failed to create the server: %v", err)
	}

	a, err := handler.NewAlice(daprAddress)
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
