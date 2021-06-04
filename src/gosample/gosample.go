package gosample

import (
	"fmt"
)

var Message string = "hello world"

type User struct {
	FirstName string
	LastName string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User {
		FirstName: firstName,
		LastName: lastName,
	}
}

type Task5 struct {
	ID int
	Detail string
	done bool
	*User
}

func NewTask(id int, detail, firstName, lastName string) *Task5 {
	task := &Task5 {
		ID: id,
		Detail: detail,
		done: false,
		User: NewUser(firstName, lastName),
	}
	return task
}

func Execute() {
	fmt.Println("execute ")

	// Task5に埋め込まれたUserのフォールドやメソッドは、Task5が実装しているか
	// のように振る舞う
	task := NewTask(1, "hogehoge", "Jxck", "Daniel")
	fmt.Println(task.FirstName)
	fmt.Println(task.LastName)
	fmt.Println(task.FullName())
	fmt.Println(task.User)

	// 型の変換
	var i uint8 = 3
	var j uint32 = uint32(i)
	fmt.Println("j = ", j)

	var s string = "abc"
	var b []byte = []byte(s)
	fmt.Println("byte = ", b)

	// Type Assertion
	Print("abc")
	Print(12345)

	Print2("abcde")
	Print2(00000)

}

func Print(value interface{}) {
	s, ok := value.(string)
	if ok {
		fmt.Println("value is string: ", s)
	} else {
		fmt.Println("value is not string")
	}
	
}

func Print2(value interface{}) {

	switch v:= value.(type) {
	case string:
		fmt.Println("value is string : ", v)
	case int:
		fmt.Println("value is int : ", v)
	case Stringer:
		fmt.Println("value is Stringer")
	}
}

type Stringer interface {
	String() string
}