package main

import . "fmt"

func main() {
	i := 90
	// i 값이 60이상이면 합격, 아니면 불합격 출력
	if i >= 60 {
		Println("합격")
	} else {
		Println("불합격")
	}

	// year에는 연도가 저장
	year := 2022

	// 윤년의 조건
	// 4배수이고 100의 배수는 아닌 경우, 400의 배수인 경우
	// 둘 중 하나의 조건을 만족하면 윤년입니다.
	// year가 윤년이면 윤년이라고 출력하고, 윤년이 아니라면 2월은 28이라고 출력

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		Println("2월은 29일")
	} else {
		Println("2월은 28일")
	}
}
