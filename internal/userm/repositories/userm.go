package repositories

import (
	"context"

	"github.com/glu/shopvui/internal/userm/models"
)

type UserRepository struct {
}

func (r *UserRepository) GetByEmail(ctx context.Context, db models.DBTX, email string) (*models.User, error) {
	q := models.New(db)
	result, err := q.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	return result, nil
}
