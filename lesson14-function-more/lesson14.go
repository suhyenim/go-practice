package main

import "fmt"

// Go 언어의 함수는 일급 변수입니다.
// Go 언어의 함수는 일련의 코드를 실행해주는 함수의 역할뿐만 아니라,
// 변수에 할당할 수 있고, 다른 함수의 인자나 리턴 값이 될 수도 있습니다.

func hamsu1() {
	fmt.Println("나는 함수!")
}

func hamsu2(x int) {
	fmt.Println(x)
}

func hamsu4(myHamsu func(int) int) {
	fmt.Println(myHamsu(9))
}

// Go 언어는 클로저(Closure)를 지원합니다.
// 클로저는 함수 안에서 함수를 선언 및 정의할 수 있고,
// 바깥쪽 함수에 선언된 변수에도 접근할 수 있는 함수를 말합니다.

func returnHamsu(x string) func() {
	sum := 0
	return func() {
		fmt.Println(x, sum)
	}
}

func main() {
	// 변수에 함수를 할당할 수 있음
	x := hamsu1
	x()
	y := hamsu2
	y(7)

	// 익명 함수
	hamsu3 := func(x int) int {
		return x * -1
	}
	fmt.Println(hamsu3(7))

	// 함수가 다른 함수의 인자나 리턴 값이 될 수도 있음
	hamsu4(hamsu3)

	// 클로저 사용
	returnHamsu("나는 클로저!")()
	z := returnHamsu("너도 클로저!")
	z()
}
