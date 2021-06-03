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
}