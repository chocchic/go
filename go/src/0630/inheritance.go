package main

import . "fmt"

type super struct {
	num  int
	name string
}

// super 구조체에 연결된 메서드
func (s *super) method() {
	Println(("super의 method"))
}

type embedding struct {
	// 구조체 안에 다른 구조체 데이터가 필드로 존재하면 has a 관계
	s     super
	score int
}

type sub struct {
	// 구조체 타입 이름만 적는 경우는 is a 관계 - 객체 지향 언어의 상속과 유사
	super
	score int
}
type person struct {
}

func (p *person) display() {
	Println("Person의 display")
}

type student struct {
	p person
}

// 오버라이딩 비스무리한 것 오버라이딩은 아님 상속이 아니기 때문
func (s *student) display() {
	// 포함된 구조체의 메서드 호출
	s.p.display()
	Println("student의 display")
}
func main() {
	e := new(embedding)
	// has a 관계일 때는 포함된 구조체 필드를 통해서 내부 구조체의 메서드 호출
	e.s.method()

	child := new(sub)
	child.method()

	s := new(student)
	s.display()
}
