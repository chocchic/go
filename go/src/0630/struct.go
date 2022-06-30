package main

import . "fmt"

type myStruct struct {
	// 정수 필드
	intField int
	// 문자열 필드
	stringField string
	// 슬라이스 필드
	sliceField []int
}

func main() {
	// 구조체 초기화
	var s1 = myStruct{
		intField:    1,
		stringField: "촉촉한 초코칩",
		sliceField:  []int{1, 2, 3, 4, 5},
	}
	Println(s1)
}
