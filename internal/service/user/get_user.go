package service

import (
	"context"
	"fmt"

	"github.com/nurkenti/furnitureShop/db/sqlc"
)

func GetUserByEmail(q *sqlc.Queries) (sqlc.User, error) {
	fmt.Println("Поиск пользователя")
	email, err := AddInfo("Email:")
	if err != nil {
		return sqlc.User{}, err
	}
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}
