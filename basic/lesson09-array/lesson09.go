package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int // 인덱스 0부터 시작
	arr1[0] = 69
	fmt.Println(arr1)
	fmt.Println(arr1[0])

	arr2 := [3]int{3, 4, 5}
	fmt.Println(arr2)
	fmt.Println(len(arr2)) // len(): 배열 길이
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Println(sum)

	arr3 := [3][2]int{{3, 1}, {4, 1}, {5, 3}} // 이차원 배열
	fmt.Println(arr3)
	fmt.Println(arr3[0][1])
}
