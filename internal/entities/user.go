package entities

import "github.com/jackc/pgtype"

type User struct {
	ID          pgtype.Text
	Email       pgtype.Text
	Password    pgtype.Text
	LastName    pgtype.Text
	FirstName   pgtype.Text
	Active      pgtype.Bool
	Inserted_at pgtype.Timestamptz
	Updated_at  pgtype.Timestamptz
}

func (u *User) TableName() string {
	return "users"
}
