package httpsample

import (
	"encoding/json"
	f "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var t = template.Must(template.ParseFiles("index.html"))

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"-"`
	Age int `json:"age"`
	Address string `json:"address,omitempty"`
	memo string
}

type Person2 struct {
	ID int `json:"id"`
	Name string `json:"name"`
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

	// ファイル書き込み
	hello := []byte("hello world\n")
	err4 := ioutil.WriteFile("./file3.txt", hello, 0666)
	if err != nil {
		log.Fatal(err4)
	}

	// ファイルの中身を全て読み込み
	message3, err5 := ioutil.ReadFile("./file3.txt")
	if err5 != nil {
		log.Fatal(err5)
	}
	f.Print("file3.txt = ", string(message3))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	f.Fprint(w, "hello world")

}

func PersonHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	// メソッドがPOST
	if r.Method == "POST" {

		var person Person2
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode((&person))
		if err != nil {
			log.Fatal(err)
		}

		// ファイル名を{id}.txtとする
		filename := f.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}
		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		// パラメータ取得
		id, err := strconv.Atoi((r.URL.Query().Get("id")))
		if err != nil {
			log.Fatal(err)
		}

		filename := f.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// person生成
		person := Person{
			ID: id,
			Name: string(b),
		}

		// レスポンスにエンコーディグしたHTMLを書き込む
		t.Execute(w, person)

	}
}