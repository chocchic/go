package main

import (
	. "fmt"
	r "runtime"
	t "time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		Println("Printing : ", i)
		t.Sleep(1 * t.Second)
	}
}

func main() {
	// 코어의 갯수
	Println(r.NumCPU())
	// Go가 Core의 개수만큼만 동시에 수행할 수 있는 개수를 설정해서 사용
	r.GOMAXPROCS(r.NumCPU())

	for i := 0; i < 100; i++ {
		go runSomeLoop(3)
	}
	t.Sleep(3 * t.Second)

}
