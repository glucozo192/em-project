package main

import (
	"context"
	"net"
	"os"

	userService "github.com/glu/shopvui/internal/userm/services"

	productService "github.com/glu/shopvui/internal/productmgt/services"
	"github.com/glu/shopvui/pkg/pb"
	"github.com/glu/shopvui/pkg/pb/product"
	"github.com/glu/shopvui/util"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var srv server

func main() {
	ctx := context.Background()
	if err := srv.LoadConfig(ctx); err != nil {
		return
	}
	if srv.config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if err := srv.loadDatabase(ctx); err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}
	go runProductServer(srv.config, srv.conn)
	go runRestful(srv.config, srv.conn)
	runUserServer(srv.config, srv.conn)
}
func runUserServer(config util.Config, conn *pgx.Conn) {
	grpcLogger := grpc.UnaryInterceptor(userService.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	userServer := userService.NewUserService(config, conn)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.UserPort)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot listen:")
	}

	log.Printf("start User server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server:")
	}
}

func runProductServer(config util.Config, conn *pgx.Conn) {
	grpcLogger := grpc.UnaryInterceptor(productService.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	productServer := productService.NewProductService(conn)
	product.RegisterProductServiceServer(grpcServer, productServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ProductPort)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot listen:")
	}

	log.Printf("start Product server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server:")
	}
}

func runRestful(config util.Config, conn *pgx.Conn) {
	server, err := userService.NewServer(config, conn)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server:")
	}
}
