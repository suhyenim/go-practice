package main

import (
	"fmt"
)

func main() {
	answer1 := 5 > 9 || 7 < 9
	fmt.Printf("%t\n", answer1)

	answer2 := 5 > 9 && 7 < 9
	fmt.Printf("%t\n", answer2)

	answer3 := !true
	fmt.Printf("%t\n", answer3)

	answer4 := !!true
	fmt.Printf("%t\n", answer4)

	x := 5
	answer5 := !(3 > 2 && 2 < 4) && (false && true) || x < 9
	fmt.Printf("%t\n", answer5)
}
