package services

import (
	"context"
	"fmt"

	"github.com/glu/shopvui/golibs/database"
	"github.com/glu/shopvui/internal/productmgt/entities"
	"github.com/glu/shopvui/internal/productmgt/repositories"
	pb "github.com/glu/shopvui/pkg/pb/product"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductService struct {
	//config util.Config
	pb.UnimplementedProductServiceServer
	DB database.Ext

	ProductRepo interface {
		InsertProduct(ctx context.Context, db database.Ext, e *entities.Product) (*entities.Product, error)
	}
}

func NewProductService(db database.Ext) pb.ProductServiceServer {
	return &ProductService{
		DB:          db,
		ProductRepo: new(repositories.ProductRepo),
	}
}

func toInsertProduct(req *pb.InsertProductRequest) *entities.Product {
	return &entities.Product{
		ID:            database.Text(uuid.NewString()),
		Sku:           database.Text(req.Product.Sku),
		Name:          database.Text(req.Product.Name),
		Description:   database.Text(req.Product.Description),
		RegularPrice:  database.Int4(req.Product.RegularPrice),
		DiscountPrice: database.Int4(req.Product.DiscountPrice),
		Quantity:      database.Int4(req.Product.Quantity),
		Taxable:       database.Bool(req.Product.Taxable),
	}
}

func (s *ProductService) InsertProduct(ctx context.Context, req *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {
	fmt.Println("come InsertProduct ne`")
	product := toInsertProduct(req)
	fmt.Println("check toInsertProduct()")
	p, err := s.ProductRepo.InsertProduct(ctx, s.DB, product)
	fmt.Println("check  s.ProductRepo.InsertProduct(ctx, s.DB, product)")
	if err != nil {
		return nil, err
	}
	return &pb.InsertProductResponse{
		Product: &pb.Product{
			Id:            p.ID.String,
			Sku:           p.ID.String,
			Name:          p.Name.String,
			Description:   p.Description.String,
			RegularPrice:  p.RegularPrice.Int,
			DiscountPrice: p.DiscountPrice.Int,
			Quantity:      p.Quantity.Int,
			Taxable:       p.Taxable.Bool,
			CreatedAt:     timestamppb.New(p.CreatedAt.Time),
			UpdatedAt:     timestamppb.New(p.UpdatedAt.Time),
		},
	}, nil
}
