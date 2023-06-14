package entities

import "github.com/jackc/pgtype"

type Role struct {
	ID         pgtype.Text
	Name       pgtype.Text
	InsertedAt pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

func TableName(name string) string {
	return "roles"
}
