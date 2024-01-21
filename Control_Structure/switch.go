package main

import "fmt"

func main() {
	//Switch Case
	// words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, " 	is a short word!")
	// 	case 5:
	// 		wordlen := len(word)
	// 		fmt.Println(word, " is exactly the right length: ", wordlen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, " is a long word")
	// 	}
	// }

	//Output
	// a  is a short word!
	// cow  is a short word!
	// smile  is exactly the right length:  5
	// anthropologist  is a long word

	//Switch Case with missing label
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, " is even")
		case i%3 == 0:
			fmt.Println(i, " is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("Exit the loop!")
			break
		default:
			fmt.Println(i, " is boring")
		}
	}

	//Output
	// 0  is even
	// 1  is boring
	// 2  is even
	// 3  is divisible by 3 but not 2
	// 4  is even
	// 5  is boring
	// 6  is even
	// Exit the loop!
	// 8  is even
	// 9  is divisible by 3 but not 2
}
