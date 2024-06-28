package main

import (
	"fmt"
	// "math" <- cos, tan 계산 라이브러리
)

func main() {
	// 1. 정수 계산
	var num1 int = 9
	var num2 int = 2
	answer12 := num1 % num2 // num1과 num2의 데이터 타입이 같아야 함
	fmt.Printf("%d\n", answer12)

	// 2. 실수 계산
	var num3 float32 = 9.0
	var num4 float32 = 2.0
	answer34 := num3 / num4
	fmt.Printf("%g\n", answer34)
}
