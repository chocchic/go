package main

import . "fmt"

/*
func main() {
	// 이 함수 호출이 나중에 이루어집니다.
	// main함수의 수행이 종료되기 직전에 호출합니다.
	defer lazy1() // 스택에 놓임
	defer lazy2() // lazy1위에 올라감
	Println("hello")
}

func lazy1() {
	Println("첫번째 지연함수")
}

func lazy2() {
	Println("두번째 지연함수")
}
*/

func main() {
	// Open한 후 작업을 처리하고 정리
	Open()
	defer Close()
	// 작업 처리
	Println("작업 수행")

	// 정리
	//Close() -> 위에서 defer Close() 하는 것과 같음
}

func Open() {
	Println("열어서 내용을 리턴")
}

func Close() {
	Println("정리")
}
