package main

import "fmt"

func main2() {
	// Using an index Expression in string
	var s string = "Hello World"
	var b byte = s[6]
	fmt.Println(string(b))

	//Slice Expression in string
	var s2 string = s[1:5]
	var s3 string = s[6:]
	fmt.Println(s2)
	fmt.Println(s3)

	//Using len function in a string
	fmt.Println(len(s))

	//Rune and byte conversion to string
	var r rune = 'X'
	fmt.Println(string(r))
	var by byte = 'X'
	fmt.Println(string(by))
}
