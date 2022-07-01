package main

import (
	. "fmt"
	t "time"
)

// 채널에 데이터를 저장해서 전송할 동시성 함수
// 채널의 자료형은 chan데이터의 자료형입니다.
func runLoopSeed(n int, ch chan int) {
	for i := 0; i < n; i++ {
		// 채널에 데이터를 저장해서 전송
		ch <- i
		t.Sleep(1 * t.Second)
	}
	close(ch)
}

// 채널로부터 데이터를 받아서 출력하는 동시성 함수
func runLoopReceive(ch chan int) {
	for {
		i, ok := <-ch // 채널의 데이터를 읽는 부분을 만들면 채널에 데이터가 저장될 때까지 대기상태가 됩니다.
		if !ok {
			break
		}
		Println("받은 데이터 : ", i)
	}
}

func main() {
	// 채널 생성
	myChannel := make(chan int)

	// 동시에 수행되도록 함수 호출
	go runLoopSeed(10, myChannel)
	go runLoopReceive(myChannel)

	t.Sleep(30 * t.Second)
}
