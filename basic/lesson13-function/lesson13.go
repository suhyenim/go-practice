package main

import "fmt"

// 함수는 특정 기능을 위해 만든 여러 문장을 묶어서 실행하는 코드 블럭 단위입니다.
// 기본적인 형태의 함수 선언은 "func" 함수이름 (매개변수이름 매개변수형) 반환형"입니다.
// 매개변수(parameter)와 전달인자(argument) 중에서 value값이 들어가면 전달인자입니다.

func hello() {
	fmt.Println("hello")
}

func print(x int) { // x는 매개변수(parameter)
	fmt.Println(x)
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func add2(x, y int) {
	fmt.Println(x + y)
}

func add3(x, y int) int {
	return x + y
}

func substract1(x, y int) (int, int) {
	return x + y, x - y
}

func substract2(x, y int) (w1, w2 int) {
	w1 = x + y
	w2 = x - y
	return
}

func substract3(x, y int) (w1, w2 int) {
	defer fmt.Println("반환 작업이 끝난 후 실행됩니다.")
	w1 = x + y
	w2 = x - y
	fmt.Println("반환 작업 중입니다.")
	return
}

func main() {
	hello()

	print(5) // 5는 전달인자(argument)

	add(5, 7)

	add2(5, 7)

	z := add3(5, 7)
	fmt.Println(z)

	z1, z2 := substract1(5, 7)
	fmt.Println(z1, z2)

	z1, z2 = substract2(5, 7)
	fmt.Println(z1, z2)

	z1, z2 = substract3(5, 7)
	fmt.Println(z1, z2)
}
