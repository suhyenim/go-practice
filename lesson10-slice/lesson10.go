package main

import "fmt"

func main() {

	// go에서는 배열보다 slice를 더 많이 사용
	// 그 이유는 배열보다 slice가 더 유연하기 때문 (배열: 길이가 정적 / slice: 길이가 가변적)
	// slice 구성은 lesson10-slice-image.jpg 참고

	var x [5]int = [5]int{1, 2, 3, 4, 5} // 길이를 지정하면 "배열"
	var s []int = x[1:3]                 // 길이를 지정하지 않으면 "slice"
	fmt.Println(s)
	fmt.Println(len(s)) // length: slice 시작부터 slice 끝까지의 길이 => x[1], x[2]
	fmt.Println(cap(s)) // capacity: slice 시작부터 원래 배열 끝까지의 길이 => x[1], x[2], x[3], x[4]
	fmt.Println(s[1:])

	var a []int = []int{1, 2, 3, 4, 5}
	b := append(a, 7, 9)
	fmt.Println(b)

	c := make([]int, 3, 4) // len이 3이고, cap이 4인 slice 생성 (0 초기화 함께 수행됨)
	c[0] = 10
	c[1] = 20
	c[2] = 30
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
