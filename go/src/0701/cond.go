package main

import (
	. "fmt"
	r "runtime"
	"sync"
)

func main() {
	r.GOMAXPROCS(r.NumCPU()) // CPU의 모든 코어 개수만큼 GO에서 사용

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			Println("wait begin:", n)
			cond.Wait() // 대기
			Println("wait end : ", n)
			mutex.Unlock() //잠금해제
		}(1)
	}

	for i := 0; i < 3; i++ {
		<-c
	}

	// 대기중인 고루틴 꺠우는 작업
	for i := 0; i < 3; i++ {
		mutex.Lock()
		Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
	}
}
