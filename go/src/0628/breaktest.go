package main

import . "fmt"

func main() {
	i := 0
	for i < 5 {
		if i%3 == 2 {
			// 반복문 종료
			//break
			// 더이상 아래 구문을 수행하지 않고 맨 위로 이동
			continue
		}
		Println(i)
		i++
	}
}
