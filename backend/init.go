package main

import (
	"collegedbms/db"
	"database/sql"

	_ "github.com/lib/pq"
)

// DB the man the myth the DB
var DB *db.Queries

func init() {
	a, err := sql.Open("postgres", "host=localhost port=5432 user=roz password=157 dbname=db sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	DB = db.New(a)
}
