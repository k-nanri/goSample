package main


import (
	// 任意の名前を指定すれば、パッケージ名を変更できる
	f "fmt"
	"example.com/gosample"
	"os"
	"errors"
)

func main() {
	f.Println(gosample.Message)

	var (
		a string = "aaaa"
		b = "bbbb"
		c = "cccc"
	)

	f.Println("a = ",a, "b = ", b, "c = ",c )
	const Hello string = "hello"
	f.Println("Hello = ", Hello)
	var i int = 0
	for {
		// 無限ループ
		i = i + 1
		if i > 10 {
			f.Println("Break")
			break
		}
	}

	file, err := os.Open("hoge.go")
	if err != nil {
		f.Println("Error")
		f.Println(err)
		f.Println(file)
	}

	n, err := dev(10, 0)
	if err != nil {
		f.Println(err)
	}
	f.Println("n = ", n)

	// 無名関数
	func(i, j int) {
		f.Println(i + j)
	}(2, 4)

	var month map[int]string = map[int]string{}
	month[1] = "Jaunary"
	month[2] = "February"
	f.Println(month)

	_, ok := month[3]
	if ok {
		f.Println("OK")
	} else {
		f.Println("NG")
	}

}

func dev(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return
	}
	result = i / j
	return
}