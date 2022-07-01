package main

import (
	. "fmt"
	r "runtime"
	"sync"
	t "time"
)

func main() {
	r.GOMAXPROCS(r.NumCPU()) // CPU의 모든 코어 개수만큼 GO에서 사용

	// 공유 데이터 생성
	var data = []int{} // 정수형 슬라이스 생성

	// 공유 자원을 동시에 사용할 수 없도록 해주는 Mutex 생성
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			// Unlock()이 호출될 때까지 이 영역의 자원은 다른 곳에서 사용할 수 없음
			mutex.Lock()
			data = append(data, 1) // 1을 1000번 집어넣기
			mutex.Unlock()
		}
	}() // 익명함수 바로 실행

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1) // 1을 1000번 집어넣기
			mutex.Unlock()
		}
	}() // 익명함수 바로 실행

	t.Sleep(2 * t.Second)
	Println(len(data)) // 슬라이스의 길이 출력
}
