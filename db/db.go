package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func New() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@tcp(localhost:3306)/sekigae?charset=utf8mb4&parseTime=true")
}
