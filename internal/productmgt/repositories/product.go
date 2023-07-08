package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/glu/shopvui/golibs/database"
	"github.com/glu/shopvui/internal/productmgt/entities"
	"go.uber.org/multierr"
)

type ProductRepo struct {
}

func (p *ProductRepo) InsertProduct(ctx context.Context, db database.Ext, e *entities.Product) (*entities.Product, error) {
	fmt.Println("check come to InsertProduct")
	fieldNames, values := database.FieldMap(e)
	placeHolders := database.GeneratePlaceholders(len(fieldNames))
	query := fmt.Sprintf(`
		INSERT INTO products (%s) VALUES (%s)
		ON CONFLICT ON CONSTRAINT pk_products DO UPDATE SET
		name = excluded.name
		`, strings.Join(fieldNames, ", "), placeHolders)
	//fmt.Println(query)
	now := time.Now()
	err := multierr.Combine(
		e.CreatedAt.Set(now),
		e.UpdatedAt.Set(now),
	)
	if err != nil {
		return &entities.Product{}, fmt.Errorf("multierr.Combine: %w", err)
	}
	_, err = db.Exec(ctx, query, values...)
	if err != nil {
		return &entities.Product{}, fmt.Errorf("db.Exec: %w", err)
	}
	return e, nil
}
