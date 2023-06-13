package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/glu/shopvui/internal/entities"
	"github.com/glu/shopvui/internal/golibs/database"
	"github.com/jackc/pgtype"
)

type UserRepo struct {
}

const getUserStmTpl = `
	select %s from %s where email = $1
`

func (r *UserRepo) GetUser(ctx context.Context, db database.QueryExecer, email pgtype.Text) (*entities.User, error) {
	e := entities.User{}
	fields, values := database.FieldMap(e)
	err := db.QueryRow(ctx, fmt.Sprintf(getUserStmTpl, strings.Join(fields, ","), "users"), &email).Scan(values)
	if err != nil {
		return nil, fmt.Errorf("db.QueryRow.Scan: %w", err)
	}
	return &e, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, db database.QueryExecer, u *entities.User) (*entities.User, error) {

	return u, nil
}
