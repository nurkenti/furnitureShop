package service

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nurkenti/furnitureShop/db/sqlc"
)

func Authorisation(q *sqlc.Queries) error {
	fmt.Println("Регистрация")
	var em string
	name, err := AddInfo("Name: ")
	if err != nil {
		fmt.Println("Name is not correct")
		return err
	}
	for {
		email, err := AddInfo("Email: ")
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
	pw, err := AddInfo("Password: ")
	if err != nil {
		return err
	}
	age, err := AddInfo("Age: ")
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
		return fmt.Errorf("Error with Create %e", err)
	}

	FormatInfo(user)
	return nil
}
func AddInfo(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin) //Читает целую строку до \n
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	return "", scanner.Err()
}

func FormatInfo(user sqlc.User) {
	fmt.Printf("Name:%s \nEmail:%s\n", user.FullName, user.Email)
	fmt.Printf("   ID: %s\n", uuid.UUID(user.ID.Bytes).String())
	fmt.Printf("   Возраст: %d\n   Роль: %s\n", user.Age, user.Role.UserRole)
	fmt.Printf("   Создан: %s\n", user.CreatedAt.Time.Format("2006-01-02 15:04:05"))
}
