package main

import (
	"bufio" // 입력 관련
	"fmt"
	"os"      // 운영체제 관련
	"strconv" // 문자열 전환 관련
)

func main() {
	// #1
	fmt.Printf("무언가 말씀해주세요!\n")
	scan1 := bufio.NewScanner(os.Stdin)
	scan1.Scan()
	input1 := scan1.Text()
	fmt.Printf("당신이 쓴 글은: %q\n", input1)

	// #2
	fmt.Printf("나이가 몇살이세요?\n")
	scan2 := bufio.NewScanner(os.Stdin)
	scan2.Scan()
	input2, _ := strconv.ParseInt(scan2.Text(), 10, 64)
	fmt.Printf("당신의 나이는 100살까지 %d년 남았습니다.\n", 100-input2)
}
