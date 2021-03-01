package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/dapr/components-contrib/metadata"
	daprc "github.com/dapr/go-sdk/client"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	alice "github.com/kzmake/dapr-kit/example/binding-priority-queue/api/alice/v1"
	bob "github.com/kzmake/dapr-kit/example/binding-priority-queue/event/bob/v1"
)

// NewHandled ...
func NewHandled(name string, in *alice.HandleRequest, md map[string]string) (*daprc.InvokeBindingRequest, error) {
	fmt.Printf("in: %+v", in)
	fmt.Printf("md: %+v", md)

	if md == nil {
		md = map[string]string{}
	}

	e := cloudevents.NewEvent(cloudevents.VersionV1)
	e.SetID(uuid.New().String())
	e.SetSource("example/uri")
	e.SetType("example.type")

	if err := e.SetData(cloudevents.ApplicationJSON, in); err != nil {
		return nil, err
	}

	d, err := e.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return &daprc.InvokeBindingRequest{
		Name:      name,
		Operation: "create",
		Metadata:  md,
		Data:      d,
	}, nil
}

// Alice ...
type Alice struct {
	alice.UnimplementedAliceServiceServer

	name string
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

	e := &bob.HandledRequest{
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
