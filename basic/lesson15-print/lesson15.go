package main

import "fmt"

// 참고: https://golang.org/pkg/fmt/*/

func main() {
	// Printf() 함수는 포맷을 지정해줘야 함
	fmt.Printf("안녕 나의 데이터 형은 %T, 그리고 값은 %d야\n", 10, 10)
	fmt.Printf("불리언은 %t\n", true)
	fmt.Printf("2020년은 10진수로는 %d, 16진수로는 %X\n", 2020, 2020)
	fmt.Printf("2020년은 과학적 표현으로 %e야\n", 2020.000000000)
	fmt.Printf("2020년은 과학적 표현으로 %f야\n", 2020.000000000)
	fmt.Printf("2020년은 과학적 표현으로 %g야\n", 2020.000000000)
	fmt.Printf("2020년은 %s\n", "코로나의 해로 기억된다")
	fmt.Printf("2020년은 %q\n", "코로나의 해로 기억된다")
	fmt.Printf("2020년은 %111q\n", "코로나의 해로 기억된다")
	fmt.Printf("\t %.3f년은 %q\n", 2020.12345, "코로나의 해로 기억된다")

	// Sprintf() 함수는 메모리에 출력함
	var x string = fmt.Sprintf("%T %d\n", 10, 10)
	print(x)
}
