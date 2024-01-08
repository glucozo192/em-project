package repositories

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/glu/shopvui/internal/userm/models"
	"github.com/jackc/pgxutil"
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

func (r *UserRepository) Create(ctx context.Context, db models.DBTX, user *models.User) error {
	q := models.New(db)

	b, err := json.Marshal(user)
	if err != nil {
		return err
	}

	var params models.CreateUserParams
	if err := json.Unmarshal(b, &params); err != nil {
		return err
	}
	_, err = q.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateUserV2(ctx context.Context, db models.DBTX, user *models.User) error {
	var params models.CreateUserParams
	m := mapCommonFields(user, params)

	_, err := pgxutil.Insert(ctx, db, "users", m)
	if err != nil {
		return err
	}
	return nil
}

func mapCommonFields(model1 any, model2 any) []map[string]any {
	result := make([]map[string]interface{}, 0)

	val1 := reflect.ValueOf(model1)
	val2 := reflect.ValueOf(model2)

	typ1 := reflect.TypeOf(model1)
	typ2 := reflect.TypeOf(model2)

	for i := 0; i < val1.NumField(); i++ {
		field1 := typ1.Field(i)
		tag1 := field1.Tag.Get("json")

		for j := 0; j < val2.NumField(); j++ {
			field2 := typ2.Field(j)
			tag2 := field2.Tag.Get("json")

			if tag1 == tag2 {
				data := make(map[string]interface{})
				data[tag1] = val1.Field(i).Interface()
				data[tag2] = val2.Field(j).Interface()
				result = append(result, data)
			}
		}
	}

	return result
}
