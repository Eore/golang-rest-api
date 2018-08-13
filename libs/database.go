package libs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER = "postgres"
	DB_PASS = ""
	DB_NAME = "golang"
)

func Connection() *sql.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, _ := sql.Open("postgres", dbURL)
	return db
}
