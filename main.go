package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nurkenti/furnitureShop/db"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	service "github.com/nurkenti/furnitureShop/internal/service/user"
	"github.com/nurkenti/furnitureShop/menu"
)

type Service struct {
	queries *sqlc.Queries
}

const (
	dbSourse = "postgresql://nurken:123nura123@127.0.0.1:5433/furnitureShop?sslmode=disable"
)

func main() {
	fmt.Println("Connect with db")
	queries, conn, err := db.NewDB()
	if err != nil {
		log.Fatal("❌ Ошибка подключения к БД:", err)
	}
	/*conn, err := pgx.Connect(context.Background(), dbSourse)
	if err != nil {
		log.Fatal("sex")
	}
	defer conn.Close(context.Background())
	Queries = sqlc.New(conn)
	*/
	defer conn.Close(context.Background()) // Закрываем соединение
	fmt.Println("✅ Подключение к БД успешно!")

	server := &Service{
		queries: queries,
	}

	server.MenuLogin()
}

func (s *Service) GetUser() {
	for {
		user, err := service.GetUserByEmail(s.queries)
		if err != nil {
			fmt.Println("Not found this email\nPlease try again\n---")
			continue
		}
		service.FormatInfo(user)
		break
	}
}

func Salesman() {

	menu.Doing()
}
func (s *Service) DeleteUserByEm() {
	err := service.DeleteUserByEmail(s.queries)
	if err != nil {
		log.Fatal("Cannot delete the user")
	}
}

func (s *Service) ListUsers() {
	err := service.ListUsers(s.queries, 5, 5)
	if err != nil {
		log.Fatal("Error with list (")
	}

}
func (s *Service) AuthorisationUser() {
	err := service.Authorisation(s.queries)
	if err != nil {
		log.Fatal("Error")
	}
}
func (s *Service) MenuLogin() {
	fmt.Println("Welcome")
	for {
		ans, err := service.AddInfo("1.Authorisation\n2.Get user\n3.List\n4.Delete\nexit\nPlease write number: ")
		if err != nil {
			log.Fatal("Your answer is not correct")
		}
		if strings.Contains(ans, "1") {
			s.AuthorisationUser()
		}
		if strings.Contains(ans, "2") {
			s.GetUser()
		}
		if strings.Contains(ans, "3") {
			s.ListUsers()
		}
		if strings.Contains(ans, "4") {
			s.DeleteUserByEm()
		}
		if ans == "exit" {
			break
		}
	}
}
