// Code generated by protoc-gen-gotemplate. DO NOT EDIT.
package v1

import (
	context "context"
	fmt "fmt"

	daprc "github.com/dapr/go-sdk/client"
	common "github.com/dapr/go-sdk/service/common"
	invoke "github.com/kzmake/dapr-kit/invoke"
	grpc "google.golang.org/grpc"
)

const AliceServiceName = "api.alice.v1.AliceService"

// aliases
type (
	invocationHandler = func(context.Context, *common.InvocationEvent) (*common.Content, error)
)

// AliceServiceHandler ...
type AliceServiceHandler interface {
	Handle(context.Context, *HandleRequest) (*HandleResponse, error)
}

// RegisterAliceServiceInvocationHandlers ...
func RegisterAliceServiceInvocationHandlers(s common.Service, impl AliceServiceHandler) error {
	return invoke.RegisterInvocationHandlers(s, impl, AliceServiceName)
}

type AliceServiceInvocationClient interface {
	Handle(context.Context, *HandleRequest, ...grpc.CallOption) (*HandleResponse, error)
}

type AliceserviceInvocationClient struct {
	appID string
	conn  *grpc.ClientConn
}

func NewAliceServiceInvocationClient(appID string, conn *grpc.ClientConn) AliceServiceInvocationClient {
	return &AliceserviceInvocationClient{
		appID: appID,
		conn:  conn,
	}
}

func (c *AliceserviceInvocationClient) Handle(ctx context.Context, in *HandleRequest, opts ...grpc.CallOption) (*HandleResponse, error) {
	cc := daprc.NewClientWithConnection(c.conn)

	reqData, err := invoke.Marshal(in)
	if err != nil {
		return nil, err
	}

	resData, err := cc.InvokeMethodWithContent(ctx, c.appID, fmt.Sprintf("%s.%s", AliceServiceName, "Handle"), "POST", &daprc.DataContent{
		ContentType: invoke.ContentType,
		Data:        reqData,
	})
	if err != nil {
		return nil, err
	}

	out := &HandleResponse{}
	if err := invoke.Unmarshal(resData, out); err != nil {
		return nil, err
	}

	return out, nil
}
