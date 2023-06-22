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
}

func NewProductService(config util.Config, db database.Ext) product.ProductServiceServer {
	return &ProductService{
		config: config,
		DB:     db,
	}
}
