// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {

// 	//Shadowing variables
// 	x := 10
// 	if x > 5 {
// 		fmt.Println(x)
// 		x := 5
// 		fmt.Println(x)
// 	}
// 	fmt.Println(x)

// 	//Output
// 	// 10
// 	// 5
// 	// 10

// 	//Shadowing with multiple assignment
// 	x:=10
// 	if x>5 {
// 		x,y := 5,20
// 		fmt.Println(x,y)
// 	}
// 	fmt.Println(x)

// 	//Output
// 	//5 20
// 	//10

// 	//Shadowing package names
// 	x := 10
// 	fmt.Println(x)
// 	fmt := "oops"
// 	fmt.Println(fmt)

// 	//Output
// 	//fmt.Println undefined (type string has no field or method Println)

// 	//if and else
// 	n := rand.Intn(10)
// 	if n == 0 {
// 		fmt.Println("That's too low.")
// 	} else if n > 5 {
// 		fmt.Println("That's too big :", n)
// 	} else {
// 		fmt.Println("That's a perfect number :", n)
// 	}

// 	//Scoping a variable to an if
// 	if n := rand.Intn(10); n == 0 {
// 		fmt.Println("That's too low.")
// 	} else if n > 5 {
// 		fmt.Println("That's too big :", n)
// 	} else {
// 		fmt.Println("That's a perfect number :", n)
// 	}
// 	fmt.Println(n)

// 	//Output
// 	//undefined: n
// }
