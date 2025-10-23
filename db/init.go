package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	db "github.com/nurkenti/furnitureShop/db/sqlc"
)

const (
	dbSource = "postgresql://nurken:123nura123@127.0.0.1:5433/furnitureShop?sslmode=disable"
)

func NewDB() (*db.Queries, *pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		return nil, nil, err
	}
	return db.New(conn), conn, nil
}
