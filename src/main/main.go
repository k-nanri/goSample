package main


import (
	// 任意の名前を指定すれば、パッケージ名を変更できる
	f "fmt"
	"example.com/gosample"
)

func main() {
	f.Println(gosample.Message)
}