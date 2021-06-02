package main


import (
	// 任意の名前を指定すれば、パッケージ名を変更できる
	f "fmt"
	"example.com/gosample"
	"os"
	"errors"
)

type ID int
type Priority int

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

	var id ID = 123
	var priority Priority = 567
	ProcessTask(id, priority)

	var task Task = Task {
		ID: 1,
		Detail: "buy the milk",
		done: true,
	}

	f.Println("ID = ", task.ID)
	f.Println("Detail = ", task.Detail)
	f.Println("done = ", task.done)

	// 初期化してない場合はゼロで初期化される
	task = Task{}
	f.Println("ID = ", task.ID)
	f.Println("Detail = ", task.Detail)
	f.Println("done = ", task.done)


	task2 := &Task{done: false}
	Finish(task2)
	f.Println(task2.done)

	task3 := NewTask(3, "hoge")
	f.Printf("%+v", task3)

	task5 := NewTask(5, "hoge2")
	task5.Finish2()
	f.Printf("%+v", task5)

}

func (task *Task) Finish2() {
	task.done = true
}

// コンストラクタはNewで始まる関数を定義
func NewTask(id int, detail string) *Task {
	task := &Task{
		ID: id,
		Detail: detail,
		done: false,
	}
	return task
}

type Task struct {
	ID int
	Detail string
	done bool
}

func Finish(task *Task) {
	task.done = true
}

func ProcessTask(id ID, priority Priority) {
	f.Println("id = ", id, ", priority = ", priority)
}

func dev(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return
	}
	result = i / j
	return
}