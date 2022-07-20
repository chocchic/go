package main

import (
	"fmt"
)

func main() {
	states := []string{"A", "B", "C", "D", "E"}
	graph := []int{1, 2, 3, -2, 4}
	n, temp := 0, 23

	for {
		old_temp := temp
		temp += graph[n]
		state := n
		if temp > 36 {
			temp = old_temp
			temp += graph[3]
			state = 3
		}
		n = (n + 1) % len(graph)

		fmt.Println("현재 온도 : ", temp, "state : ", states[state])
	}

}
