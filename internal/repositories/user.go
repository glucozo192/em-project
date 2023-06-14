package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/glu/shopvui/internal/entities"
	"github.com/glu/shopvui/internal/golibs/database"
	"github.com/jackc/pgtype"
	"go.uber.org/multierr"
)

type UserRepo struct {
}

const getUserStmTpl = `
	select %s from %s where email = $1
`

func (r *UserRepo) GetUser(ctx context.Context, db database.Ext, email pgtype.Text) (*entities.User, error) {
	e := &entities.User{}
	fields, values := database.FieldMap(e)
	f, v := e.FieldMap()
	fmt.Println("f and v: ", f, v)
	fmt.Println("fields and value: ", fields, values)
	err := db.QueryRow(ctx, fmt.Sprintf(getUserStmTpl, strings.Join(fields, ","), "users"), email).Scan(values...)
	if err != nil {
		return nil, fmt.Errorf("db.QueryRow.Scan: %w", err)
	}
	fmt.Println(values...)
	return e, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, db database.Ext, e *entities.User) (*entities.User, error) {
	fieldNames, values := database.FieldMap(e)
	placeHolders := database.GeneratePlaceholders(len(fieldNames))
	query := fmt.Sprintf(`
		INSERT INTO users (%s) VALUES (%s)
		ON CONFLICT ON CONSTRAINT pk_users DO UPDATE SET
			first_name = excluded.first_name,
			last_name = excluded.last_name,
            password = excluded.password,
			updated_at = excluded.updated_at
		`, strings.Join(fieldNames, ", "), placeHolders)
	now := time.Now()
	err := multierr.Combine(
		e.InsertedAt.Set(now),
		e.UpdatedAt.Set(now),
		e.Active.Set(true),
	)
	if err != nil {
		return &entities.User{}, fmt.Errorf("multierr.Combine: %w", err)
	}
	fmt.Println(e)
	_, err = db.Exec(ctx, query, values...)
	if err != nil {
		return &entities.User{}, fmt.Errorf("db.Exec: %w", err)
	}
	return e, nil
}

// func (r *UserRepo) AddRoles(ctx context.Context, db database.Ext, roles *entities.Role) error {
// 	query:= `INSERT INTO roles (%s) VALUES (%s)`
// 	fields, values := roles.FieldMap()
// 	return nil
// }
