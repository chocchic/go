package main

import . "fmt"

// call by value 함수 - 매개변수가 참조형이 아닌 경우
func callByValue(n int) {
	n = n + 1
	Println("n : ", n)
}

// call by reference 함수 - 매개변수가 참조형(포인터)인 경우
func callByReference(p *int) {
	*p = *p + 1
	Println("*p :", *p)
}

func main() {
	var numPtr *int
	Println(numPtr)
	Printf("numPtr : %d\n", numPtr)

	// 동적 메모리 공간을 할당한 후 값을 대입하고 값을 출력
	numPtr = new(int)
	*numPtr = 30
	Println(numPtr, " : ", *numPtr)

	// 다른 메모리 영역의 주소 가리키기
	x := 123
	numPtr = &x
	Println(numPtr, ":", *numPtr)

	// 더블 포인터
	doublePointer := &numPtr
	Println(doublePointer, ":", *doublePointer)

	x = 1
	// x는 value 타입이라서 함수의 매개변수로 대입해도 값이 변경되지 않습니다.
	callByValue(x)
	Println("x: ", x)

	// x의 참조를 넘기기 때문에 데이터가 변경될 수도 있습니다.
	// 참조를 넘겼는데 데이터의 변경이 안되는 경우는 그 함수가 데이터 변경자체를 안하는 경우인데, 이 경우는 리턴이 반드시 있습니다.
	callByReference(&x)
	Println("x: ", x)
}
