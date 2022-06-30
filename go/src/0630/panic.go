package main

import . "fmt"

// 패닉을 확인하는 함수
func checkPanic() {
	// 패닉이 발생하면 프로그램을 중단하지 않고 패닉 메세지를 출력
	if r := recover(); r != nil {
		Println("Panic Caputred", r)
	}
}

// panic을 발생시키는 함수
func panictest(p bool) {
	defer checkPanic()
	if p == true {
		panic("문제 발생!! 프로그램을 종료합니다.")
	}
}

func main() {
	panictest(true)
}
