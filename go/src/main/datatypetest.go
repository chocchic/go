package main

import . "fmt"

func main() {
	var x int = 20
	Println(x)

	// 선언과 동시에 초기화하는 경우는 자료형 생략이 가능
	var y = 30
	Println(y)

	// 서로 다른 자료형의 데이터를 동시에 초기화
	var i, f = 10, 20.7
	Printf("x=%d y=%f", i, f)

	//var str = "Hello World"
	str := "Hello World" // 위와 결과 같음
	Printf("str=%s\n", str)
}
