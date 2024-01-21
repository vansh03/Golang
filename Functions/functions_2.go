package main

import "fmt"

//Variadic imput parameters and slices
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, []int{1, 2, 3, 4}...))
	fmt.Println(addTo(3, 1, 2, 3, 4))
}
