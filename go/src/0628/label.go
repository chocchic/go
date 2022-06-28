package main

import . "fmt"

func main() {
Outer: // 다음에 나타나는 제어문에 이름을 붙이는 것으로 이름 뒤에 :을 해주어야 합니다.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Outer // 바깥 쪽 반복문 빠져나감
			}
			Println(i, j)
		}
	}
}
