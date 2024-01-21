package main

import "fmt"

func main3() {

	//Declaring a empty map using var declaration
	var nilMap map[string]int
	fmt.Println(len(nilMap))

	//Declaring a map using map literal
	literalMap := map[string]int{}
	fmt.Println(literalMap)
	fmt.Println(len(literalMap))
	fmt.Println(nil == literalMap)

	//nonempty map lietral

	teams := map[string][]string{
		"Mi":  []string{"Fred", "Bread"},
		"Csk": []string{"Dhoni", "Poni"},
	}
	fmt.Println(teams)
	fmt.Println(len(teams))

	//Known key=value pair
	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	fmt.Println(len(ages))

	//Reading and writing a map
	totalWins := map[string]int{}
	totalWins["Vansh"] = 3
	totalWins["Bread"] = 1
	fmt.Println(totalWins["Vansh"])
	fmt.Println(totalWins["Fred"])
	totalWins["Fred"]++
	fmt.Println(totalWins["Fred"])

	//comma ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 3,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	a, ok := m["world"]
	fmt.Println(a, ok)

	a, ok = m["Bie"]
	fmt.Println(a, ok)

	//Deleting from maps
	deleteMap := map[string]int{
		"hello": 5,
		"world": 2,
	}
	delete(deleteMap, "hello")
	fmt.Println(deleteMap)

	//Using map as Sets
	intSet := map[int]bool{}
	vals := []int{1, 3, 3, 5, 7, 5, 8, 9, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[10])
	fmt.Println(intSet[100])

}
