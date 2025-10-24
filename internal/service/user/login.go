package service

import (
	"context"
	"fmt"

	"github.com/nurkenti/furnitureShop/db/sqlc"
)

func Login(q *sqlc.Queries) error {
	fmt.Println("Login")
	for {
		email, err := AddInfo("Email: ")
		if err != nil {
			return err
		}
		pw, err := AddInfo("Password: ")
		if err != nil {
			return err
		}
		user, err := q.GetUserByEmail(context.Background(), email)
		if err != nil {
			fmt.Println("[Email has not found] \nPlease try again")
			continue

		}
		if pw != user.PasswordHash {
			fmt.Println("[Your pw has not correct] \nPlease try again")
			continue
		} else {
			fmt.Printf("Welcome %s\n", user.FullName)
			break
		}
	}
	return nil
}
