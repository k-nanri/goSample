package main

import (
	"time"
	"fmt"
)
//import "local/mypkg"



func main() {
	fmt.Println("Hello, World")

	var a1 int
	a1 = 1
	fmt.Println("a1 = ", a1)

	var s1 string
	s1 = "hoge"
	fmt.Println("s1 = ", s1)

	var s2 string = "aaabbbccc"
	fmt.Println("s2 = ", s2)

	// type
	type Name string
	var firstName Name = "Tanaka"
	fmt.Println("First Name = ", firstName)

	// array
	var aa1 [3]string
	aa1[0] = "Red"
	aa1[1] = "Green"
	aa1[2] = "Blue"
	fmt.Println(aa1[0], aa1[1], aa1[2])


	// スライス == Javaでいうリスト？
	aa2 := []string{}
	aa2 = append(aa2, "Red")
	aa2 = append(aa2, "Blue")
	fmt.Println(aa2[0], aa2[1])
	
	// capは容量。メモリ上に確保されている数。
	aa3 := []int{}
	for i := 0; i< 10; i++ {
		aa3 = append(aa3, i)
		fmt.Println(len(aa3), cap(aa3))
	}

	// メモリの事前確保
	//bufa := make([]byte, 0 , 1024)

	// map
	map1 := map[string]int{
		"x": 100,
		"y": 200,
	}

	fmt.Println(map1["x"])

	// map add value
	map1["z"] = 300
	fmt.Println(map1["z"])

	delete(map1, "z")

	// mapをループ処理する
	fmt.Println("### Map Loop ####")
	for key,value := range map1 {
		fmt.Println(key ," = ", value)
	}

	var x int = 3
	var y int = 4

	if x < y {
		fmt.Println("True")
	}

	var mode string = "stop"

	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}

	// go では switch文にbreakなし. 処理を継続する場合はfallthroughを指定

	var dayOfWeek string = "sat"
	switch dayOfWeek {
	case "sat":
		fmt.Println("in case")
		fallthrough
	case "sun":
		fmt.Println("Horiday")
	default:
		fmt.Println("Weekday")
	}

	// for
	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}

	// cotinue, break
	n := 0
	for {
		n++
		if n > 10 {
			fmt.Println("Break!!")
			break
		} else if n % 2 == 1 {
			fmt.Println("Continue = ", n)
			continue
		} else {
			fmt.Println(n)
		}
	}

	// rangeを使ってloop
	colors := [...]string{"Red", "Green", "Blue"}
	for i, color := range colors {
		fmt.Println(i, ": ", color)
	}

	// goto
	fmt.Println("goto1")
	goto Done
	fmt.Println("goto2")
	Done:
	fmt.Println("Done")

	// function
	var z int = add(3, 5)
	// z := add(3, 5)
	fmt.Println("add function return values is ", z)

	// function retuns multiple values
	var x4, y4 int = addMinus(10, 20)
	fmt.Println("x4 = ", x4)
	fmt.Println("y4 = ", y4)

	// 不要な戻り値がある場合は、ブランク変数 _ を使用することができる
	_, y4 = addMinus(10, 30)
	fmt.Println("y4 = ", y4)

	// ...を用いると可変引数を実現可能
	fmt.Println("Call funcA !!!!")
	funcA(1, 2, 3, 4, 5, 6, 7)

	// struct
	var p1 Person
	p1.SetPerson("Yamada", 26)
	name, age := p1.GetPerson()
	fmt.Println("Person name = ", name, ", age = ", age)

	var p2 Person2
	p2.SetPerson2("Tanaka")
	p2.Age = 30
	fmt.Println("Person2 Age = ", p2.Age)

	fmt.Println("=== interface check =================")
	person3 := Person{name: "takaki"}
	book3 := Book{title: "hogehoge"}
	PrintOut(person3)
	PrintOut(book3)

	fmt.Println("=== interface{} ==================")
	funcB(book3)

	// pointer
	fmt.Println("=== Pointer ===================")
	var pointer1 int
	var pointer2 *int

	pointer2 = &pointer1
	*pointer2 = 123
	fmt.Println(pointer1)
	var b1 int = 123
	var b2 int = 123
	pointercheck(b1, &b2)
	fmt.Println("b1 = ", b1, ", b2 = ", b2)

	var pi3 Person
	pi3.SetPerson("takahashi", 27)
	p3 := &pi3
	fmt.Println("pi3.name = ", pi3.name)
	fmt.Println("(*p3).name = ", (*p3).name)
	fmt.Println("p3.name = ", p3.name)

	// new

	bookList := []*Book{}

	for i:= 0; i < 10; i++ {
		book := new(Book)
		book.title = fmt.Sprintf("Title#%d", i)
		bookList = append(bookList, book)
	}

	for _, book := range bookList {
		fmt.Println(book.title)
	}

	go funcGoroutine()
	for i := 0; i < 20; i++ {
		fmt.Println("M")
		time.Sleep(200 * time.Millisecond)
	}



}

func funcGoroutine() {
	fmt.Println("funcGoroutine start")
	for i := 0; i < 10; i++ {

		fmt.Println("A")
		time.Sleep(100 * time.Millisecond)
	}
}



func pointercheck(b1 int, b2 *int) {
	b1 = 456
	*b2 = 789
}



func funcA(a int, b ... int) {
	fmt.Println("a = ", a)
	for i, num := range b {
		fmt.Println("b[",i,"] = ", num)
	}	
}

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

func add(x int, y int) int {
	return x + y
}

type Person struct {
	name string
	age int
}

type Person2 struct {
	name string
	Age int
}

func (p *Person2) SetPerson2(name string) {
	p.name = name
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name ,p.age
}

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) SetBook(title string) {
	b.title = title
}

func (b Book) ToString() string {
	return b.title
}

func funcB(a interface{}) {
	q, ok := a.(Printable)
	if ok {
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not Printable.")
	}
}


