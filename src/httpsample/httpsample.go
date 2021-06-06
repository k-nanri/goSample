package httpsample

import (
	f "fmt"
	"encoding/json"
	"log"
)

type Person struct {
	ID int
	Name string
	Email string
	Age int
	Address string
	memo string
}



func ExecuteHttp() {
	f.Println("=== ExecuteHttp ===================")

	person := &Person {
		ID: 1,
		Name: "Gopher",
		Email: "gopher@eample.com",
		Age:10,
		Address:"Japan",
		memo: "golang study",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	f.Println(string(b))

	
}