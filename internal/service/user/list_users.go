package service

import (
	"context"
	"fmt"

	"github.com/nurkenti/furnitureShop/db/sqlc"
)

func ListUsers(q *sqlc.Queries, limit, offs int32) error {
	fmt.Println()
	arg := sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offs,
	}
	users, err := q.ListUsers(context.Background(), arg)
	if err != nil {
		return err
	}
	for _, user := range users {
		FormatInfo(user)
	}
	return nil
}
