package main

import "fmt"

func main() {
	var number1 int = 29      // 자료형 명시적 선언 (1)
	number2 := 2900000.788888 // 자료형 명시적 선언 (2)
	fmt.Printf("%T %T\n", number1, number2)

	var number3 = 29 // 자료형 묵시적 선언
	fmt.Printf("%T\n", number3)

	var number4 uint64 // 기본값: 0
	var boo bool       // 기본값: false
	fmt.Println(number4, boo)

	// 주의: 한 번 자료형을 선언하고 나면, 변경 불가능
}
