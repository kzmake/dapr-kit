package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	bob "github.com/kzmake/dapr-kit/example/request-replay/api/bob/v1"

	"github.com/kzmake/dapr-kit/example/request-replay/bob/handler"
)

const (
	appAddress = ":4000"
)

func main() {
	s, err := daprd.NewService(appAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	b := &handler.Bob{}
	if err := bob.RegisterBobServiceInvocationHandler(s, b); err != nil {
		log.Fatalf("failed to register invocation handlers: %v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
