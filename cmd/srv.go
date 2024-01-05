package cmd

import (
	"context"

	"github.com/glu/shopvui/idl/pb"
	postgres_client "github.com/glu/shopvui/pkg/postgres"

	"github.com/glu/shopvui/configs"
	"github.com/glu/shopvui/utils"
	"github.com/jackc/pgx/v4"
)

var srv server

type processor interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type factory interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}

type server struct {
	cfg *configs.Config

	processors []processor
	factories  []factory

	// // repository
	// userRepo    userRepo.UserRepo
	// productRepo productRepo.ProductRepo

	// // services
	// productSrv product.ProductServiceServer
	userService pb.UserServiceServer

	// config
	config utils.Config

	// deliveries
	userDelivery pb.UserServiceServer

	// database clients
	postgresClient *postgres_client.PostgresClient
	conn           *pgx.Conn

	// processors []processor
	// factories  []factory
}

// type processor interface {
// 	Init(ctx context.Context) error
// 	Start(ctx context.Context) error
// 	Stop(ctx context.Context) error
// }

// type factory interface {
// 	Connect(ctx context.Context) error
// 	Stop(ctx context.Context) error
// }

func (s *server) loadGrpcClients(ctx context.Context) error {
	//connections
	//s.userConnClient = grpc_client.NewGrpcClient(s.cfg.UserServiceEndpoint)
	return nil
}

func (s *server) LoadConfig(ctx context.Context) error {
	var err error
	s.config, err = utils.LoadConfig(".")
	if err != nil {
		return err
	}
	return nil
}

func (s *server) loadDeliveries(ctx context.Context) error {
	//s.userDelivery = deliveries.NewUserDelivery(s.userService)
	// s.productSrv = productSrv.NewProductService(s.conn)
	return nil
}

func (s *server) loadDefault(ctx context.Context) {
	if err := s.loadDeliveries(ctx); err != nil {
		panic(err)
	}
}

func (s *server) loadPostgres(ctx context.Context) {
	var err error
	srv.postgresClient, err = postgres_client.NewClient(srv.cfg.PostgresUrl)
	if err != nil {
		panic(err)
	}
}

func start(ctx context.Context, errChan chan error) error {
	for _, f := range srv.factories {
		if err := f.Connect(ctx); err != nil {
			errChan <- err
		}
	}
	for _, p := range srv.processors {
		go func(p processor) {
			if err := p.Start(ctx); err != nil {
				errChan <- err
			}
		}(p)
	}

	return nil
}

func stop(ctx context.Context) error {
	for _, processor := range srv.processors {
		if err := processor.Stop(ctx); err != nil {
			return err
		}
	}

	for _, database := range srv.factories {
		if err := database.Close(ctx); err != nil {
			return err
		}
	}
	return nil
}
