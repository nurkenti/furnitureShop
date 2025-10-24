package service

import (
	"context"
	"fmt"
	"log"

	"github.com/nurkenti/furnitureShop/db/sqlc"
)

func DeleteUserByEmail(q *sqlc.Queries) error {
	email, err := AddInfo("Option: Delete\nEmail: ")
	if err != nil {
		return err
	}
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		return err
	}
	FormatInfo(user)
	err = q.DeleteUserByEmail(context.Background(), email)
	if err != nil {
		log.Fatal("User has not delete")
	} else {
		fmt.Println("User has been delete")
	}
	return nil
}
