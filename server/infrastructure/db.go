package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@/gosample")
	if err != nil {
		return nil, err
	}
	return db, nil
}
