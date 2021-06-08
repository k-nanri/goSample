package httpsample

import (
	"encoding/json"
	f "fmt"
	"log"
	"os"
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

	// ファイルへの書き込み
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	message := []byte("hello world\n")

	// Write()で書き込み
	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	file2, err2 := os.Open("./file2.txt")
	if err2 != nil {
		f.Println("Open Error")
		log.Fatal(err2)
	}

	defer file2.Close()

	message2 := make([]byte, 20)

	_, err2 = file2.Read(message2)
	if err2 != nil {
		
		log.Fatal(err2)
	}

	f.Println("file2.txt = " , string(message2))
}