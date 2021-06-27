package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
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
}

func slowHandler(w http.ResponseWriter, req *http.Response) {
	start := time.Now()
	err := slowQuery()
	if err != nil {
		log.Printf("Error : %s\n", err.Error())
		return
	}
	fmt.Fprintln(w, "OK")
	fmt.Printf("slowHandler took: %v\n", time.Since(start))
}

func main() {

	var err error
	db, err = sql.Open("postgres", conn)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	srv := http.Server {
		Addr: ":8080",
		WriteTimeout:2 * time.Second,
		Handler: http.HandlerFunc(slowHandler),
	}

	log.Println("STart Http Server...")
	log.Fatal(srv.ListenAndServe())
}