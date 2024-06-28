package main

import "fmt"

func main() {
	// 1. for
	x := 1
	for x < 5 {
		fmt.Println(x)
		x++
	}
	for x := 1; x < 5; x += 2 {
		fmt.Println(x)
	}

	// 2. continue, break
	for x := 1; x <= 15; x++ {
		if x == 3 {
			continue
		} else if x == 8 {
			break
		}
		fmt.Println(x)
	}

}
