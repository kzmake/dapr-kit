package handler

import (
	"context"
	"log"
	"time"

	bob "github.com/kzmake/dapr-kit/example/binding-priority-queue/event/bob/v1"
)

// Bob ...
type Bob struct {
	bob.UnimplementedBobServiceServer
}

// interfaces
var _ bob.BobServiceServer = (*Bob)(nil)
var _ bob.BobServiceHandler = (*Bob)(nil)

// Handle ...
func (c *Bob) Handle(_ context.Context, req *bob.HandledRequest) (*bob.HandledResponse, error) {
	log.Printf("[%s] recv *bob.HandleRequest: %s", req.GetId(), req.String())
	log.Printf("[%s] \t do something...", req.GetId())
	time.Sleep(10 * time.Second)
	log.Printf("[%s] ...completed", req.GetId())
	return &bob.HandledResponse{}, nil
}
