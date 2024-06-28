package main

import (
	"fmt"
)

func main() {
	x := 7
	y := 3
	answer1 := x < y // x와 y의 데이터 타입이 같아야 함
	fmt.Printf("%t\n", answer1)

	a := "a"
	b := "b"
	answer2 := a < b // 아스키값으로 판단
	fmt.Printf("%t\n", answer2)
}
