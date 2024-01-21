package main

import "fmt"

//declaring a struct
type person struct {
	name string
	age  int
	pet  string
}

func main1() {

	fmt.Println(person{"bob", 20, "cat"})
	fmt.Println(person{name: "Fred"})

	//struct literal can be map literal style
	bob := person{
		age:  30,
		name: "vansh",
	}
	fmt.Println(bob.name)

	//struct literal can be comma seperated
	julie := person{
		"jilie",
		40,
		"cat",
	}
	fmt.Println(julie.age)
}
