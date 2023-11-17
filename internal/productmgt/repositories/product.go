package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/glu/shopvui/internal/productmgt/entities"
	"github.com/glu/shopvui/internal/userm/golibs/database"
	"github.com/jackc/pgtype"
	"go.uber.org/multierr"
)

type ProductRepo struct {
}

func (p *ProductRepo) InsertProduct(ctx context.Context, db database.Ext, e *entities.Product) (*entities.Product, error) {
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

func (p *ProductRepo) GetProductByID(ctx context.Context, db database.Ext, id pgtype.Text) (*entities.Product, error) {
	e := &entities.Product{}
	fieldNames, values := database.FieldMap(e)
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE product_id = $1`, strings.Join(fieldNames, ", "), e.TableName())
	err := db.QueryRow(ctx, query, id).Scan(values...)
	if err != nil {
		return e, err
	}
	return e, nil
}

func (p *ProductRepo) ListAllProducts(ctx context.Context, db database.Ext) ([]*entities.Product, error) {
	e := make([]*entities.Product, 0)
	product := &entities.Product{}
	fieldNames, _ := database.FieldMap(product)
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE inserted_at is not null`, strings.Join(fieldNames, ", "), product.TableName())

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("DB exc: %v", err)
	}
	for rows.Next() {
		pTemp := &entities.Product{}
		_, values := database.FieldMap(pTemp)
		if err := rows.Scan(values...); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		e = append(e, pTemp)
	}
	return e, nil
}
