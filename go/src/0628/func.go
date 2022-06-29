package main

import . "fmt"

// 함수 선언
func display() {
	for i := 0; i < 10; i++ {
		Println(i)
	}
	Println("함수 종료")
}

func main() {
	// function call
	display()
	Println("동잃ㄴ")
	display()
}
