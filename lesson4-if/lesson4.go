package main

import "fmt"

func main() {
	// 1. if
	name := "hello"
	fmt.Println("if 이전")
	if name == "hello" {
		fmt.Println("if 안")
	}
	fmt.Println("if 밖")

	// 2. if - else
	age := 24
	if age >= 20 {
		fmt.Println("You're an university student.")
	} else {
		fmt.Println("You're not an school student.")
	}

	// 3. if - else if - else
	age = 16
	if age >= 20 {
		fmt.Println("You're an university student.")
	} else if age >= 17 {
		fmt.Println("You're a high school student.")
	} else {
		fmt.Println("You're a middle school student.")
	}

	// 4. switch
	age = 8
	switch age {
	case 8, 9, 10:
		fmt.Println("초등 저학년")
	case 11, 12, 13:
		fmt.Println("초등 고학년")
	default:
		fmt.Println("초등 아님")
	}
	age = 13
	switch {
	case age >= 8 && age <= 10:
		fmt.Println("초등 저학년")
	case age >= 11 && age <= 13:
		fmt.Println("초등 고학년")
	default:
		fmt.Println("초등 아님")
	}
}
