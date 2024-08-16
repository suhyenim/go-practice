package main

import "fmt"

func main() {
	// var 맵이름 map[key자료형]value자료형

	// #1
	var zoo map[string]int = map[string]int{
		"코끼리": 31,
		"사자":  7,
		"호랑이": 6, // 마지막도 쉼표로 끝내기
	}
	zoo["코알라"] = 90        // kv쌍 추가
	delete(zoo, "코알라")     // kv쌍 삭제
	fmt.Println(zoo["사자"]) // 해당 key에 대한 value값 출력
	val, ok := zoo["킹콩"]   // 해당 key에 대한 value값이 존재하는지 true or false로 return
	fmt.Println(val, ok)
	fmt.Println(len(zoo))
	fmt.Println(zoo)

	// #2
	aqua := make(map[string]int) // make로 빈 map 생성
	aqua["돌고래"] = 11
	aqua["상어"] = 5
	aqua["펭귄"] = 24
	fmt.Println(aqua)
}
