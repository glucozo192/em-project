package repositories

import (
	"context"
	"testing"

	"github.com/glu/shopvui/internal/entities"
	"github.com/glu/shopvui/internal/golibs/database"
	"github.com/glu/shopvui/mocks/testutil"
	"github.com/glu/shopvui/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func UserRepoWithSqlMock() (*UserRepo, *testutil.MockDB) {
	r := &UserRepo{}
	return r, testutil.NewMockDB()
}

func createRandomUser(t *testing.T) entities.User {
	ctx := context.Background()
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	userRepo, MockDB := UserRepoWithSqlMock()

	arg := &entities.User{
		ID:        database.Text(uuid.NewString()),
		Email:     database.Text(util.RandomOwner()),
		Password:  database.Text(hashedPassword),
		FirstName: database.Text(util.RandomOwner()),
		LastName:  database.Text(util.RandomEmail()),
	}

	user, err := userRepo.CreateUser(ctx, MockDB.DB, arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.NotZero(t, user.InsertedAt)

	return *user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
