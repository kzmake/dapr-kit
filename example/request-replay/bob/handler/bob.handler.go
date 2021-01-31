package handler

import (
	"context"

	bob "github.com/kzmake/dapr-kit/example/request-replay/api/bob/v1"
)

// Bob ...
type Bob struct {
	bob.UnimplementedBobServiceServer
}

// interfaces
var _ bob.BobServiceServer = (*Bob)(nil)
var _ bob.BobServiceHandler = (*Bob)(nil)

// Handle ...
func (c *Bob) Handle(_ context.Context, req *bob.HandleRequest) (*bob.HandleResponse, error) {
	return &bob.HandleResponse{Msg: "called Bob"}, nil
}
