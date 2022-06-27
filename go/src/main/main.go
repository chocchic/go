package main

import "fmt"

func main() {
	fmt.Println("Hello go")
	fmt.Println("Hello")
	fmt.Println("go")

	fmt.Print("Hello")
	fmt.Print("go")

	fmt.Printf("i:%d", 20)
	fmt.Printf("f:%f\n", 20.65231)
	// 소수 3번째 자리까지만 출력 - 4번째 자리에서 반올림
	fmt.Printf("f:%.3f\n", 20.64123)
	// 각 정수를 5자리씩 자리를 확보한 후 출력
	fmt.Printf("i: %5d%5d\n", 20, 10)
	// 앞의 빈자리에 0을 추가
	fmt.Printf("i:%05d%05d", 20, 10)

	j := 10
	if j >= 5 {
		fmt.Println("5이상의 숫자")
	} else {
		fmt.Println("5미만의 숫자")
	}
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	// 한줄에 2개의 명령문이 존재하는 경우는 각 명령문을 구분하기 위해서 ;를 사용
	fmt.Println("하나의 출력 명령어")
	fmt.Println("하나의 출력 명령어2")
}
