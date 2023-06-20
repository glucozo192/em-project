package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/glu/shopvui/internal/services"
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
		log.Fatal("cannot load config:", err)
	}
	conn, err := pgx.Connect(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	go runGatewayServer(config, conn)
	runGrpcServer(config, conn)
}

func runRestful(config util.Config, conn *pgx.Conn) {
	server, err := services.NewServer(config, conn)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runGrpcServer(config util.Config, conn *pgx.Conn) {
	grpcServer := grpc.NewServer()
	userServer := services.NewUserService(config, conn)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	log.Printf("start gPRC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGatewayServer(config util.Config, conn *pgx.Conn) {
	userServer := services.NewUserService(config, conn)

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
		log.Fatal("cannot register gRPC handler server:", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	log.Printf("start gPRC server at %s", listener.Addr().String())

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start gateway server:", err)
	}
}
