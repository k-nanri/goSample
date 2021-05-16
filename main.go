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

}