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

type Rect struct {
	width  int
	height int
}

// Rect 구조체의 필드의 값을 가져오는 메서드와 설정하는 메서드
func (rect *Rect) getWidth() int {
	return rect.width
}
func (rect *Rect) getHeight() int {
	return rect.height
}
func (rect *Rect) setWidth(width int) {
	rect.width = width
}
func (rect *Rect) setHeight(height int) {
	rect.height = height
}

// 메서드를 만들 때 *을 생략하고 리시버를 만들면 리시버는 호출하는 리시버의 데이터가 새로 만들어져서 대입됩니다.
func (rect Rect) copyHeight(height int) {
	rect.height = height
}

// 리시버가 메서드 안에서 사용이 안될 때는 _로 이름을 만드는 것이 가능
func (_ *Rect) display() {
	Println("나는 길이와 너비를 가지고 있습니다.")
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

	// Rect 구조체와 연결된 메서드 사용
	rect := Rect{}
	rect.setHeight(10)
	rect.setWidth(20)
	Println(rect.getWidth(), rect.getHeight())

	// copyHeight 메서드는 height를 매개변수의 값으로 변경하도록 되어있지만
	// 리시버가 포인터 형태가 아니기 때문에 호출하는 리시버의 데이터를 변경할 수 없습니다.
	rect.copyHeight(30)
	Println(rect.getHeight(), rect.getHeight())

	rect.display()
}
