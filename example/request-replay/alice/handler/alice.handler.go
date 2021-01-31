package handler

import (
	"context"
	"log"

	"google.golang.org/grpc"

	alice "github.com/kzmake/dapr-kit/example/request-replay/api/alice/v1"
	bob "github.com/kzmake/dapr-kit/example/request-replay/api/bob/v1"
)

// Alice ...
type Alice struct {
	alice.UnimplementedAliceServiceServer

	bobc bob.BobServiceInvocationClient
}

// interfaces
var _ alice.AliceServiceServer = (*Alice)(nil)
var _ alice.AliceServiceHandler = (*Alice)(nil)

// NewAlice ...
func NewAlice(bobAppID, address string) (alice.AliceServiceServer, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	bobc := bob.NewBobServiceInvocationClient(bobAppID, conn)

	return &Alice{bobc: bobc}, nil
}

// Handle ...
func (c *Alice) Handle(ctx context.Context, req *alice.HandleRequest) (*alice.HandleResponse, error) {
	res, err := c.bobc.Handle(ctx, &bob.HandleRequest{})
	if err != nil {
		log.Printf("failed: %v", err)
		return nil, err
	}

	return &alice.HandleResponse{Msg: "called Alice" + ": " + res.GetMsg()}, nil
}
