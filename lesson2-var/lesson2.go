package main

import "fmt"

func main() {
	// var 변수명 자료형
	// 변수명: 맨 앞 숫자 X, 공백 X, 특수문자 X, 예약어 X
	// 자료형: int, bool, string 등

	var test1 string       // 변수 선언
	test1 = "Hello world2" // 변수 초기화
	test1 = "Hello world3"
	fmt.Println(test1)

	var number uint8 = 29 // 변수 선언 및 초기화
	number = number + 100
	fmt.Println("Hello world!", number)
}
