package grpc

import (
	"net"

	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	dapr "github.com/dapr/go-sdk/service/grpc"
	"google.golang.org/grpc"
)

// Service ...
type Service struct {
	dapr.Server
	listener   net.Listener
	grpcServer *grpc.Server
}

// New ...
func New(s *grpc.Server, lis net.Listener) common.Service {
	return &Service{
		grpcServer: s,
		listener:   lis,
	}
}

// Start ...
func (s *Service) Start() error {
	pb.RegisterAppCallbackServer(s.grpcServer, s)
	return s.grpcServer.Serve(s.listener)
}

// Stop ...
func (s *Service) Stop() error {
	s.grpcServer.GracefulStop()

	return nil
}
