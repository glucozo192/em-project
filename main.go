package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	userService "github.com/glu/shopvui/internal/userm/services"
	"github.com/glu/shopvui/pkg/pb"
	"github.com/glu/shopvui/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config:")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	conn, err := pgx.Connect(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	go runGatewayServer(config, conn)
	runGrpcServer(config, conn)
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

func runGrpcServer(config util.Config, conn *pgx.Conn) {
	grpcLogger := grpc.UnaryInterceptor(userService.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	userServer := userService.NewUserService(config, conn)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot listen:")
	}

	log.Printf("start gPRC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server:")
	}
}

//func runUnaryServer(grpcServer *grpc.Server) error

func runGatewayServer(config util.Config, conn *pgx.Conn) {
	userServer := userService.NewUserService(config, conn)

	// userServer := services.NewUserService(conn)
	// pb.RegisterUserServiceServer(grpcServer, userServer)
	// reflection.Register(grpcServer)

	// config outputs
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userServer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register gRPC handler server:")
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot listen:")
	}

	log.Printf("start gPRC server at %s", listener.Addr().String())

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gateway server:")
	}
}
