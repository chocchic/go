package main

import (
	. "fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

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

	result5 := RecursiveSum(7)
	Println(result5)

	result6 := Fib(10)
	Println(result6)

	result7 := FibnoRecursive(10)
	Println(result7)

	// 함수를 변수에 저장
	f := F1
	f()

	F2()

	// 익명함수 생성과 호출
	r := func(n int) int {
		if n == 1 {
			return 1
		}
		return n + RecursiveSum(n-1)
	}(5)
	Println(r)

	// 파일 내용을 읽어오기
	// 첫번째 리턴되는 내용은 데이터를 읽은 경우 데이터
	// 두번째 리터되는 내용은 에러가 발생하지 않으면 nil, 에러가 발생하면 에러 내용
	b, err := ioutil.ReadFile("./file.txt")
	if err == nil {
		Printf("contents:%s\n", b)
	}

	if b, err := ioutil.ReadFile("./file.txt"); err == nil {
		Printf("contents:%s\n", b)
	}

	// switch 표현식에서 함수 호출
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(10); {
	case i >= 1 && i < 6:
		Println("작은 숫자")
	default:
		Println("큰 숫자")
	}

	// outer함수에 만들어진 변수 접근 - 에러
	/*
		localvar = localvar + 1
		Println(localvar)
	*/

	// Outer가 리턴한 함수를 대입
	clousre := outer()
	clousre()
	outer()

}

// 변수에 함수를 할당해서 생성
var F2 = func() {
	Println("함수")
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

// 1부터 n까지의 합을 리턴하는 함수 - n+ n-1까지의 합
func RecursiveSum(n int) int {
	if n == 1 {
		return 1
	}

	return n + RecursiveSum(n-1)
}

// 피보나치 수열 - 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// 첫번째와 두번째는 1
// 세번째부터는 이전 2개 항의 합

func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func FibnoRecursive(n int) int {
	a := 1
	b := 1
	result := 1
	for i := 3; i <= n; i++ {
		result = a + b
		a = b
		b = result
	}
	return result
}

func outer() func() {
	localvar := 1

	// 리턴하는 함수 - 이 함수는 Outer안에 존재하기 때문에 localVar을 사용하는 것이 가능
	return func() {
		localvar = localvar + 1
		Println("localvar : ", localvar)
	}
}
