// Code generated by protoc-gen-gotemplate. DO NOT EDIT.
package v1

import (
	context "context"
	fmt "fmt"

	daprc "github.com/dapr/go-sdk/client"
	common "github.com/dapr/go-sdk/service/common"
	binding "github.com/kzmake/dapr-kit/binding"
	encoding "github.com/kzmake/dapr-kit/encoding"
	proto "github.com/kzmake/dapr-kit/encoding/proto"
	invocation "github.com/kzmake/dapr-kit/invocation"
	grpc "google.golang.org/grpc"
)

const _ = daprc.UndefinedType
const _ = common.SubscriptionResponseStatusSuccess
const _ = grpc.SupportPackageIsVersion7

var _ fmt.Stringer
var _ context.Context
var _ invocation.Def
var _ binding.Def

const BobServiceName = "event.bob.v1.BobService"

// BobServiceHandler ...
type BobServiceHandler interface {
	Handle(context.Context, *HandleRequest) (*HandleResponse, error)
}

// RegisterBobServiceInvocationHandler ...
func RegisterBobServiceInvocationHandler(s common.Service, impl BobServiceHandler) error {
	fns := map[string]invocation.Func{
		"event.bob.v1.BobService/Handle": _BobService_Handle_Invocation_Handler(impl.Handle),
	}

	for name, fn := range fns {
		if err := s.AddServiceInvocationHandler(name, fn); err != nil {
			return err
		}
	}

	return nil
}

func _BobService_Handle_Invocation_Handler(handler interface{}) invocation.Func {
	return func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
		e, err := encoding.New(in.ContentType)
		if err != nil {
			return nil, err
		}

		req := new(HandleRequest)
		if err := e.Unmarshal(in.Data, req); err != nil {
			return nil, err
		}

		fn := handler.(func(context.Context, *HandleRequest) (*HandleResponse, error))
		res, err := fn(ctx, req)
		if err != nil {
			return nil, err
		}

		d, err := e.Marshal(res)
		if err != nil {
			return nil, err
		}

		out := &common.Content{
			DataTypeURL: "event.bob.v1.BobService/Handle",
			ContentType: "application/json",
			Data:        d,
		}

		return out, nil
	}
}

// RegisterBobServiceBindingHandler ...
func RegisterBobServiceBindingHandler(s common.Service, impl BobServiceHandler) error {
	fns := map[string]binding.Func{
		"event.bob.v1.BobService/Handle": _BobService_Handle_Binding_Handler(impl.Handle),
	}

	for name, fn := range fns {
		if err := s.AddBindingInvocationHandler(name, fn); err != nil {
			return err
		}
	}

	return nil
}

func _BobService_Handle_Binding_Handler(handler interface{}) binding.Func {
	return func(ctx context.Context, in *common.BindingEvent) ([]byte, error) {
		e := proto.NewEncoding()

		req := new(HandleRequest)
		if err := e.Unmarshal(in.Data, req); err != nil {
			return nil, err
		}

		fn := handler.(func(context.Context, *HandleRequest) (*HandleResponse, error))
		res, err := fn(ctx, req)
		if err != nil {
			return nil, err
		}

		d, err := e.Marshal(res)
		if err != nil {
			return nil, err
		}

		return d, nil
	}
}

type BobServiceInvocationClient interface {
	Handle(context.Context, *HandleRequest, ...grpc.CallOption) (*HandleResponse, error)
}

type bobserviceInvocationClient struct {
	appID string
	conn  *grpc.ClientConn
}

func NewBobServiceInvocationClient(appID string, conn *grpc.ClientConn) BobServiceInvocationClient {
	return &bobserviceInvocationClient{
		appID: appID,
		conn:  conn,
	}
}

func (c *bobserviceInvocationClient) Handle(
	ctx context.Context, in *HandleRequest, opts ...grpc.CallOption,
) (*HandleResponse, error) {
	cc := daprc.NewClientWithConnection(c.conn)

	e := proto.NewEncoding()

	req, err := e.Marshal(in)
	if err != nil {
		return nil, err
	}

	res, err := cc.InvokeMethodWithContent(ctx, c.appID, "event.bob.v1.BobService/Handle", "POST", &daprc.DataContent{
		ContentType: "application/x-protobuf",
		Data:        req,
	})
	if err != nil {
		return nil, err
	}

	out := new(HandleResponse)
	if err := e.Unmarshal(res, out); err != nil {
		return nil, err
	}

	return out, nil
}

type BobServiceBindingClient interface {
	Handle(context.Context, *HandleRequest, map[string]string) (*HandleResponse, error)
}

type bobserviceBindingClient struct {
	conn *grpc.ClientConn
}

func NewBobServiceBindingClient(conn *grpc.ClientConn) BobServiceBindingClient {
	return &bobserviceBindingClient{
		conn: conn,
	}
}

func (c *bobserviceBindingClient) Handle(
	ctx context.Context, in *HandleRequest, meta map[string]string,
) (*HandleResponse, error) {
	cc := daprc.NewClientWithConnection(c.conn)

	e := proto.NewEncoding()

	d, err := e.Marshal(in)
	if err != nil {
		return nil, err
	}

	req := &daprc.InvokeBindingRequest{
		Name:      "event.bob.v1.BobService/Handle",
		Operation: "create",
		Data:      d,
		Metadata:  map[string]string{},
	}

	res, err := cc.InvokeBinding(ctx, req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	out := new(HandleResponse)
	if err := e.Unmarshal(res.Data, out); err != nil {
		return nil, err
	}

	return out, nil
}

type BobServiceOutputBindingClient interface {
	Handle(context.Context, *HandleRequest, map[string]string) error
}

type bobserviceOutputBindingClient struct {
	conn *grpc.ClientConn
}

func NewBobServiceOutputBindingClient(conn *grpc.ClientConn) BobServiceOutputBindingClient {
	return &bobserviceOutputBindingClient{
		conn: conn,
	}
}

func (c *bobserviceOutputBindingClient) Handle(
	ctx context.Context, in *HandleRequest, meta map[string]string,
) error {
	cc := daprc.NewClientWithConnection(c.conn)

	e := proto.NewEncoding()

	d, err := e.Marshal(in)
	if err != nil {
		return err
	}

	req := &daprc.InvokeBindingRequest{
		Name:      "event.bob.v1.BobService/Handle",
		Operation: "create",
		Data:      d,
		Metadata:  map[string]string{},
	}

	if err := cc.InvokeOutputBinding(ctx, req); err != nil {
		return err
	}

	return nil
}
