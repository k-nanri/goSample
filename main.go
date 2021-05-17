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

}