package main

import (
	. "fmt"
	t "time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		Println("Printing : ", i)
		t.Sleep(1 * t.Second)
	}
}

func main() {
	// 다른 함수와 동시에 실행되도록 앞에 go를 추가
	go runSomeLoop(10)
	go runSomeLoop(10)
	// 동시성으로 실행하는 경우에 main이 종료되지 않아야합니다.
	// main이 종료되지 않도록 하기 위해 30초 대기
	t.Sleep(30 * t.Second) // 이 줄이 없으면 스레드와 메인이 함께 끝나기 때문에 위에서 출력한 값이 보이지 않음
	// go를 붙이고부터 두개가 같이 돌아감.

}
