package main

import (
	"fmt"
	. "fmt"
)

type myStruct struct {
	// 정수 필드
	intField int
	// 문자열 필드
	stringField string
	// 슬라이스 필드
	sliceField []int
}

// 3개의 매개변수를 받아서 myStruct 구조체를 초기화한 후 리턴하는 함수
func NewMyStruct(intField int, stringField string, sliceField []int) *myStruct {
	// 스택에 만들고 그 참조를 리턴
	//return &myStruct{intField, stringField, sliceField}

	imsi := new(myStruct)
	imsi.intField = intField
	imsi.stringField = stringField
	imsi.sliceField = sliceField
	return imsi
}

func main() {
	// 구조체 초기화
	var s1 = myStruct{
		intField:    1,
		stringField: "촉촉한 초코칩",
		sliceField:  []int{1, 2, 3, 4, 5},
	}
	Println(s1)

	// 필드 단위로 접근
	Println(s1.sliceField)

	//필드 이름을 생략하고 초기화
	var s2 = myStruct{2, "축축한 초코칩", []int{10, 20, 30, 40, 50}}
	Println(s2) // 전체 출력

	// 구조체 생성을 하고 나중에 초기화
	var s3 = myStruct{}
	s3.intField = 3
	fmt.Println(s3)

	ins := new(myStruct)
	Println(ins)
	ins.intField = 20
	Println(ins) // 이렇게 출력하면 앞에 &가 붙음

	m := NewMyStruct(4, "zlzlzl", []int{102, 103, 104, 105})
	Println(m)
}
