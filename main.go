package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nurkenti/furnitureShop/db"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	product "github.com/nurkenti/furnitureShop/internal/service/product"
	user "github.com/nurkenti/furnitureShop/internal/service/user"
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
	defer conn.Close(context.Background()) // Закрываем соединение
	fmt.Print("✅ Подключение к БД успешно!\n")

	server := &Service{
		queries: queries,
	}

	/*server.MenuLogin()*/
	server.ChairMenu()
	//fmt.Println(server)
	//product.AddChair(queries)
	//product.GetChair(queries)
	//product.DeleteChair(queries)
	//product.ListChair(queries)
	//product.UpdateChair(queries)
}

func (s *Service) GetUser() {
	for {
		users, err := user.GetUserByEmail(s.queries)
		if err != nil {
			fmt.Println("Not found this email\nPlease try again\n---")
			continue
		}
		user.FormatInfo(users)
		break
	}
}

func Salesman() {

	menu.Doing()
}
func (s *Service) DeleteUserByEm() {
	err := user.DeleteUserByEmail(s.queries)
	if err != nil {
		log.Fatal("Cannot delete the user")
	}
}

func (s *Service) ListUsers() {
	err := user.ListUsers(s.queries, 5, 5)
	if err != nil {
		log.Fatal("Error with list (")
	}

}
func (s *Service) AuthorisationUser() {
	err := user.Authorisation(s.queries)
	if err != nil {
		log.Fatal("Error")
	}
}

func (s *Service) LogIn() {
	err := user.Login(s.queries)
	if err != nil {
		log.Fatal("Error with Login")
	}

}

func (s *Service) MenuLogin() {
	fmt.Println("Welcome")
	for {
		ans, err := user.AddInfo("1.Authorisation\n2.Get user\n3.List\n4.Delete\n5.Login\nexit\nPlease write number: ")
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
		if strings.Contains(ans, "5") {
			s.LogIn()
		}
		if ans == "exit" {
			break
		}
	}
}

func (s *Service) ChairMenu() {
	for {
		fmt.Println("Chair settings")
		_, i, err := menu.NewMenuTemplate("Select settings:", []string{"Add chair", "Get chair", "Delete chair", "List chairs", "Update chair", "end"})
		if err != nil {
			log.Fatal(err)
		}
		if i == 5 {
			break
		}
		if i == 0 {
			product.AddChair(s.queries)
			menu.Timeloading(1, "Loading")
		}
		if i == 1 {
			product.GetChair(s.queries)
			menu.Timeloading(1, "Loading")
		}
		if i == 2 {
			err := product.DeleteChair(s.queries)
			if err != nil {
				fmt.Print("❌Chair has not been delete\n")
			} else {
				fmt.Print("✅Chair has been delete\n")
			}
			menu.Timeloading(1, "Loading")
		}
		if i == 3 {
			product.ListChair(s.queries)
			menu.Timeloading(1, "Loading")
		}
		if i == 4 {
			product.UpdateChair(s.queries)
			menu.Timeloading(1, "Loading")
		}
	}
}
