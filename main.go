package main

import "fmt"

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

}