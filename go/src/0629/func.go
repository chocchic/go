package main

import . "fmt"

func main() {
	F1()
	add(10, 30)

	// 가변 매개변수를 가진 함수 호출
	VarArgs(10, 20)
	VarArgs(10, 20, 30)

	result := ReturnFunc(100, 200)
	Println("결과 : ", result)

	// 2개의 데이터를 리턴하는 함수 호출
	result1, result2 := MultipleReturn(100, 2000)
	Println("결과 : ", result1, " : ", result2)

	result3, result4 := MultipleReturn(100, 2000)
	Println("결과 : ", result3, " : ", result4)
}

func F1() {
	Println("f1 호출")
}

// 매개변수의 자료형이 모두 일치한다면 자료형은 1번만 기재해도 됨
// func add(a int, b int)와 결과가 같음
func add(a, b int) {
	Println(a + b)
}

func VarArgs(ar ...int) {
	total := 0
	for _, value := range ar {
		total = total + value
	}
	Println(total)
}

// 리턴이 있는 함수
func ReturnFunc(a, b int) int {
	return a + b
}

// 여러개의 데이터를 리턴하는 함수
// 2개 정도까지는 이 방식이 유용하지만 서로 다른 종류의 데이터 여러개를 이런 방식으로 리턴하면 가독성이 떨어집니다.
func MultipleReturn(a, b int) (int, int) {
	return a + b, a - b
}

// NamedReturn : 리턴할 데이터의 이름을 리턴하는 방식
func NamedReturn(a, b int) (add, sub int) {
	add = a + b
	sub = a - b
	return
}
