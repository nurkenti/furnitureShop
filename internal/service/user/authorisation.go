package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nurkenti/furnitureShop/db/sqlc"
)

func Authorisation(q *sqlc.Queries) error {
	fmt.Println("Регистрация")
	var em string
	name, err := addInfo("Name: ")
	if err != nil {
		fmt.Println("Name is not correct")
		return err
	}
	for {
		email, err := addInfo("Email: ")
		if err != nil {
			return err
		}
		if !strings.Contains(email, "@") {
			fmt.Println("Email is not correct. Please try again")
			continue
		}
		em = email
		break
	}
	pw, err := addInfo("Password: ")
	if err != nil {
		return err
	}
	age, err := addInfo("Age: ")
	if err != nil {
		return err
	}
	Age, err := strconv.Atoi(age)
	if err != nil {
		return err
	}

	arg := sqlc.CreateUserParams{
		ID:           pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Email:        em,
		PasswordHash: pw,
		FullName:     name,
		Age:          int32(Age),
		Role:         sqlc.NullUserRole{UserRole: "admin", Valid: true},
	}
	user, err := q.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	FormatInfo(user)
	return nil
}
func addInfo(prompt string) (string, error) {
	var ans string
	fmt.Print(prompt)
	_, err := fmt.Scan(&ans)
	if err != nil {
		return "", err
	}
	return ans, nil
}

func FormatInfo(user sqlc.User) {
	fmt.Printf("Name:%s \nEmail:%s\n", user.FullName, user.Email)
	fmt.Printf("   ID: %s\n", uuid.UUID(user.ID.Bytes).String())
	fmt.Printf("   Возраст: %d\nРоль: %s\n", user.Age, user.Role.UserRole)
	fmt.Printf("   Создан: %s\n\n", user.CreatedAt.Time.Format("2006-01-02 15:04:05"))
}
