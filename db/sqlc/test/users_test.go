package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	"github.com/nurkenti/furnitureShop/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) sqlc.User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := sqlc.CreateUserParams{
		ID:           pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Email:        util.RandomEmail(),
		PasswordHash: hashedPassword,
		FullName:     util.RandomName(),
		Age:          int32(util.RandomAge()),
		Role:         sqlc.NullUserRole{UserRole: "admin", Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Age, user.Age)
	require.Equal(t, arg.Role, user.Role)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreatUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := sqlc.UpdateUserParams{
		ID:       user1.ID,
		FullName: user1.FullName,
		Age:      user1.Age,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err) //require - требовать , equal - равный, coloumn - столбец
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)

	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user1.UpdateAt.Time, user2.UpdateAt.Time, time.Second)
}

func TestDeleteUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)

	// Проверяем удалился ли аккаунт
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.Error(t, err)
	require.ErrorIs(t, err, pgx.ErrNoRows) // проверяет, что ошибка err является (или оборачивает) конкретную ошибку pgx.ErrNoRows. ErrorIs для проверки конкретных типов ошибок.
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	arg := sqlc.ListUsersParams{
		Limit:  5, // сколько записей показать. (0-5)
		Offset: 0, // типо как стр. Он пропускает записи. (Если написать 5, то он будет показывать 6-10 записи)
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5) //проверяем длину
	for _, user := range users {
		require.NotEmpty(t, user) // проверяем каждую запись то что он не пустой
	}
}
