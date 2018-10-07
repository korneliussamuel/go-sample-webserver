package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Setup() *sql.DB {
	var (
		host     = os.Getenv("host")
		port     = os.Getenv("port")
		user     = os.Getenv("user")
		password = os.Getenv("password")
		dbname   = os.Getenv("dbname")
		sslmode  = os.Getenv("sslmode")
		connStr  = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode)
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
