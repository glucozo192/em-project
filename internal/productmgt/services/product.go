package services

import (
	"github.com/glu/shopvui/internal/userm/golibs/database"
	"github.com/glu/shopvui/pkg/pb/product"
	"github.com/glu/shopvui/util"
)

type ProductService struct {
	config util.Config
	product.UnimplementedProductServiceServer
	DB database.Ext

	ProductRepo interface{}
}

func NewProductService(db database.Ext) product.ProductServiceServer {
	return &ProductService{
		DB: db,
	}
}
