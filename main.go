package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/nurkenti/furnitureShop/db/sqlc"
	service "github.com/nurkenti/furnitureShop/internal/service/user"
	"github.com/nurkenti/furnitureShop/menu"
)

var Queries *sqlc.Queries

const (
	dbSourse = "postgresql://nurken:123nura123@127.0.0.1:5433/furnitureShop?sslmode=disable"
)

func main() {
	conn, err := pgx.Connect(context.Background(), dbSourse)
	if err != nil {
		log.Fatal("sex")
	}
	defer conn.Close(context.Background())
	Queries = sqlc.New(conn)

	service.Authorisation(Queries)
	//Salesman()

}

func Salesman() {
	menu.Doing()
}
