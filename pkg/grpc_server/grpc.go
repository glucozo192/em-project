package grpc_server

import (
	"context"
	"fmt"
	"net"

	//"github.com/duyledat197/go-gen-tools/config"

	"github.com/glu/shopvui/util"

	//"github.com/duyledat197/go-gen-tools/pkg/registry"
	//"github.com/duyledat197/go-gen-tools/pkg/tracing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	ServiceName string
	//Consul         *registry.ConsulRegister
	//Tracer         *tracing.TracerClient
	AuthFunction   grpc_auth.AuthFunc
	server         *grpc.Server
	Logger         *zap.Logger
	Address        util.Config
	MaxMessageSize int //* default = 0 mean 4MB
	Handlers       func(ctx context.Context, server *grpc.Server) error

	OtherOptions []grpc.ServerOption
}

func (s *GrpcServer) Init(ctx context.Context) error {
	var (
		streamInterceptors []grpc.StreamServerInterceptor
		unaryInterceptors  []grpc.UnaryServerInterceptor
		opts               []grpc.ServerOption
	)

	opts = append(opts,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamInterceptors...,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryInterceptors...,
		)))
	opts = append(opts, s.OtherOptions...)
	s.server = grpc.NewServer(
		opts...,
	)
	s.Handlers(ctx, s.server)
	return nil
}

func (s *GrpcServer) Start(ctx context.Context) error {
	_, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Address.GRPCServerAddress))
	if err != nil {
		return err
	}

	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	s.server.GracefulStop()
	return nil
}
