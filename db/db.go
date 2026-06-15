package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	var err error
	DB, err = sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=Bexmeen1111 dbname=todo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
