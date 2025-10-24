package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/nurkenti/furnitureShop/db/sqlc"
)

var testQueries *sqlc.Queries
var testDB *pgx.Conn // сделали глобал переменную

const (
	dbSourse = "postgresql://nurken:123nura123@127.0.0.1:5433/furnitureShop?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgx.Connect(context.Background(), dbSourse)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = sqlc.New(testDB)
	os.Exit(m.Run())
}
