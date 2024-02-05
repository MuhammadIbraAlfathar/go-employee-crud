package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func InitDatabase() *sql.DB {
	dsn := "root@tcp(localhost:3306)/go_employee"
	db, err := sql.Open("mysql", dsn)

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	return db
}
