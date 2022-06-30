package main

import . "fmt"

func main() {
	var numPtr *int
	Println(numPtr)
	Printf("numPtr : %d\n", numPtr)

	// 동적 메모리 공간을 할당한 후 값을 대입하고 값을 출력
	numPtr = new(int)
	*numPtr = 30
	Println(numPtr, " : ", *numPtr)

	x := 123
	numPtr = &x
	Println(numPtr, ":", *numPtr)

	doublePointer := &numPtr
	Println(doublePointer, ":", *doublePointer)
}
