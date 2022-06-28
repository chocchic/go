package main

import (
	. "fmt"
)

func main() {
	// 정수 메모리 공간을 가리킬 수 있는 포인터 변수 선언
	var iptr *int
	i := 10

	// 포인터 변수에 대한 참조를 대입
	iptr = &i
	jptr := iptr
	*iptr = 40
	Println(*jptr)
}
