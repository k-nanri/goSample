package httpsample

import (
	f "fmt"
	"encoding/json"
	"log"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"-"`
	Age int `json:"age"`
	Address string `json:"address,omitempty"`
	memo string
}



func ExecuteHttp() {
	f.Println("=== ExecuteHttp ===================")

	// JSONへの変換
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

	// JSONからの変換
	var person2 Person
	c := []byte(`{"id": 1, "name": "Gooooo", "age":5}`)
	err2 := json.Unmarshal(c, &person2)
	if err2 != nil {
		log.Fatal(err2)
	}
	f.Println("json to struct = ", person2)
	
}