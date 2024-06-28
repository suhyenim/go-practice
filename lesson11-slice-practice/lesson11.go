package main

import "fmt"

func main() {
	// range는 range(시작숫자, 종료숫자, step)의 형태로 리스트 슬라이싱과 유사합니다.
	// range의 결과는 시작숫자부터 종료숫자 바로 앞 숫자까지 컬렉션을 만듭니다.
	// 시작숫자와 step은 생략할 수 있습니다.

	// #1-1
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// #1-2
	for _, a := range a {
		fmt.Println(a)
	}

	// #1-3
	for i, v := range a {
		fmt.Printf(":%d: %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// #2-1
	var b []int = []int{1, 2, 3, 4, 5, 6, 5, 7, 8, 9}
	for i, v := range b {
		for j, v2 := range b {
			if v == v2 && j > i {
				fmt.Println(v)
			}
		}
	}

	// #2-2
	for i, v := range b {
		for j := i + 1; j < len(b); j++ {
			v2 := b[j]
			if v == v2 {
				fmt.Println(v)
			}
		}
	}
}
