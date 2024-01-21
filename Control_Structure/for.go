package main

func main() {
	//Complete for statement
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//Output
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9

	//Condition-only for statement
	// i := 1
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i = i * 2
	// }

	//Output
	// 1
	// 2
	// 4
	// 8
	// 16
	// 32
	// 64

	//Infinte for statement
	// for {
	// 	fmt.Println("Hello.")
	// }

	//Output
	// Hello.
	// Hello.
	// Hello.
	// Hello.
	// Hello.
	// Hello.
	// Hello....

	//break condition
	// i := 0
	// for {
	// 	fmt.Println("Hii")
	// 	if i > 3 {
	// 		break
	// 	}
	// 	i = i + 1
	// }

	//Output
	// Hii
	// Hii
	// Hii
	// Hii
	// Hii

	//Continue Condition
	// for i := 0; i <= 10; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("Both")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Three Only")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Five Only")
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//Output
	// Both
	// 1
	// 2
	// Three Only
	// 4
	// Five Only
	// Three Only
	// 7
	// 8
	// Three Only
	// Five Only

	//For-range loop with a slice
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for i, v := range evenVals {
	// 	fmt.Println(i, v)
	// }

	//Output
	// 0 2
	// 1 4
	// 2 6
	// 3 8
	// 4 10
	// 5 12
	//Ignoring the key in a for loop
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	fmt.Println(v)
	// }

	//Output
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12

	//When we want key not value
	// uniqueNames := map[string]bool{
	// 	"Fred": true,
	// 	"Raul": true,
	// 	"Wima": true,
	// }
	// for k := range uniqueNames {
	// 	fmt.Println(k)
	// }

	//output
	// Fred
	// Raul
	// Wima

	//Map iteration order varies
	// m := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("loop ", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	//Output
	// loop  0
	// b 2
	// c 3
	// a 1
	// loop  1
	// a 1
	// b 2
	// c 3
	// loop  2
	// a 1
	// b 2
	// c 3

	//Iterating over strings
	// samples := []string{"hello", "apple_⌘!"}

	// for _, sample := range samples {
	// 	for i, r := range sample {
	// 		fmt.Println(i, r, string(r))
	// 	}
	// }

	//Output
	// 0 104 h
	// 1 101 e
	// 2 108 l
	// 3 108 l
	// 4 111 o
	// 0 97 a
	// 1 112 p
	// 2 112 p
	// 3 108 l
	// 4 101 e
	// 5 95 _
	// 6 8984 ⌘
	// 9 33 !

	//For range value is a copy
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	//Labeling for statement
	// 	samples := []string{"hello", "apple"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	//Output
	// 0 104 h
	// 1 101 e
	// 2 108 l
	// 0 97 a
	// 1 112 p
	// 2 112 p
	// 3 108 l

}
