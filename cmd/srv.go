package cmd

import (
	"context"

	productRepo "github.com/glu/shopvui/internal/productmgt/repositories"
	productSrv "github.com/glu/shopvui/internal/productmgt/services"
	userRepo "github.com/glu/shopvui/internal/userm/repositories"
	userSrv "github.com/glu/shopvui/internal/userm/services"
	"github.com/glu/shopvui/pkg/pb"
	"github.com/glu/shopvui/pkg/pb/product"
	"github.com/glu/shopvui/util"
	"github.com/jackc/pgx/v4"
)

type server struct {
	// repository
	userRepo    userRepo.UserRepo
	productRepo productRepo.ProductRepo

	// services
	productSrv product.ProductServiceServer
	userSrv    pb.UserServiceServer

	// config
	config util.Config

	// database
	conn *pgx.Conn

	processors []processor
	factories  []factory
}

type processor interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type factory interface {
	Connect(ctx context.Context) error
	Stop(ctx context.Context) error
}

func (s *server) LoadConfig(ctx context.Context) error {
	var err error
	s.config, err = util.LoadConfig(".")
	if err != nil {
		return err
	}
	return nil
}

func (s *server) loadDatabase(ctx context.Context) error {
	var err error
	s.conn, err = pgx.Connect(ctx, s.config.DBSource)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) loadServices(ctx context.Context) error {
	//s.productSrv = productSrv.NewProductService()
	s.userSrv = userSrv.NewUserService(s.config, s.conn)
	s.productSrv = productSrv.NewProductService(s.conn)
	return nil
}

func loadServer(ctx context.Context) error {
	return nil
}
