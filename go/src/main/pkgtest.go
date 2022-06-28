package main

import . "fmt" // fmt 생략 가능
//import f "fmt"

func main() {
	// Println("하나의 출력 명령어") // import시 .로 해놨을 때
	//f.Println("fmt 패키지를 f로 가져왔을 때")
	n := 0
	var ori [100000][2]int32
	Scan(&n)

	for i := 0; i < n; i++ {
		Scan(&ori[i][0])
		Scan(&ori[i][1])
	}

	for i := 0; i < n; i++ {
		Printf("%d : %d\n", ori[i][0], ori[i][1])
	}
}
