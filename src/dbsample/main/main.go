package main

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

const (
	conn = "host=postgres port =5432 user=postgres password=secret dbname=simplebank sslmode=disable"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// slow query
func slowQuery() error {
	_, err := db.Exec("SELECT pg_sleep(5)")
	return err

func main() {

	fmt.Println("Start")
}