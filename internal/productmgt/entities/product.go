package entities

import "github.com/jackc/pgtype"

type Product struct {
	ID            pgtype.Text        `db:"product_id"`
	Sku           pgtype.Text        `db:"sku"`
	Name          pgtype.Text        `db:"name"`
	Description   pgtype.Text        `db:"description"`
	RegularPrice  pgtype.Int4        `db:"regular_price"`
	DiscountPrice pgtype.Int4        `db:"discount_price"`
	Quantity      pgtype.Int4        `db:"quantity"`
	Taxable       pgtype.Bool        `db:"taxable"`
	CreatedAt     pgtype.Timestamptz `db:"inserted_at"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at"`
}

func (u *Product) TableName() string {
	return "products"
}
