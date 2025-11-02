package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/nurkenti/furnitureShop/db/sqlc"
)

var testQueries *sqlc.Queries
var testDB *pgx.Conn // сделали глобал переменную

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
