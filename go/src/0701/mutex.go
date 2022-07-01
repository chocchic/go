package main

import (
	. "fmt"
	r "runtime"
	t "time"
)

func main() {
	r.GOMAXPROCS(r.NumCPU()) // CPU의 모든 코어 개수만큼 GO에서 사용

	// 공유 데이터 생성
	var data = []int{} // 정수형 슬라이스 생성

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) // 1을 1000번 집어넣기
		}
	}() // 익명함수 바로 실행

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) // 1을 1000번 집어넣기
		}
	}() // 익명함수 바로 실행

	t.Sleep(2 * t.Second)
	Println(len(data)) // 슬라이스의 길이 출력
}
