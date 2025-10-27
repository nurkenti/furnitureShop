package db


_ "github.com/lib/pq" 


var testQueries *Queries
var testDB *sql.DB // сделали глобал переменную

const (
	dbDriver = "postgres"
	dbSourse = "postgresql://postgres:123123@127.0.0.1:5433/simple_bank?sslmode=disable"
)
