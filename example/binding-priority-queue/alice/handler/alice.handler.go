package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/dapr/components-contrib/metadata"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	alice "github.com/kzmake/dapr-kit/example/binding-priority-queue/api/alice/v1"
	bob "github.com/kzmake/dapr-kit/example/binding-priority-queue/event/bob/v1"
)

// Alice ...
type Alice struct {
	alice.UnimplementedAliceServiceServer

	conn *grpc.ClientConn
}

// interfaces
var _ alice.AliceServiceServer = (*Alice)(nil)
var _ alice.AliceServiceHandler = (*Alice)(nil)

// NewAlice ...
func NewAlice(address string) (alice.AliceServiceServer, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Alice{
		conn: conn,
	}, nil
}

// Handle ...
func (c *Alice) Handle(ctx context.Context, req *alice.HandleRequest) (*alice.HandleResponse, error) {
	bobc := bob.NewBobServiceOutputBindingClient(c.conn)

	md := map[string]string{metadata.PriorityMetadataKey: strconv.FormatInt(req.GetPriority(), 10)}

	e := &bob.HandleRequest{
		Id:       uuid.New().String(),
		Metadata: md,
	}

	if err := bobc.Handle(ctx, e, md); err != nil {
		log.Printf("failed: %v", err)
		return nil, err
	}

	res := &alice.HandleResponse{
		Task: &alice.Task{
			Id: e.GetId(),
		},
	}

	return res, nil
}
