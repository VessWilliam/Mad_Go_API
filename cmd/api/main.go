package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != err {
		log.Fatal(err)
	}
	defer db.Close()
}
