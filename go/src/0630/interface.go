package main

import . "fmt"

type Service interface {
	GetName() string
}

type ServiceImpl struct {
	name string
}

// ServiceImpl 구조체에서 호출할 수 있는 GetName메서드를 생성
// ServiceImpl 구조체는 Service 인터페이스의 메서드를 만들었으므로 implements된 것입니다.
func (s *ServiceImpl) GetName() string {
	return s.name
}

// 인터페이스 생성 - Attack()이라는 메서드를 소유한 구조체는 Starcraft를 implements한 것임
// starcraft 변수에 구조체를 대입할 수 있습니다.
type Starcraft interface {
	Attack()
}

type Protoss struct {
	unit string
}

type Terran struct {
	unit string
}

type Zerg struct {
	unit string
}

func (t Terran) Attack() {
	Println(t.unit + " 공격")
}

func (z Zerg) Attack() {
	Println(z.unit + " 공격")
}
func (p Protoss) Attack() {
	Println(p.unit + " 공격")
}
func main() {
	// 인터페이스는 변수만 선언할 수 있고 직접 메모리 할당을 할 수 없음
	// 인터페이스는 인스턴스만 생성할 수 없습니다.
	var service Service
	Println(service)

	// 인터페이스 변수에 인터페이스를 implements한 구조체의 데이터를 대입하는 것이 가능
	service = &ServiceImpl{"서비스1"}
	// 인터페이스 타입으로 선언한 변수를 이용해서 메서드를 호출하면 대입된 구조체에 구현한 메서드가 호출됩니다.
	Println(service.GetName())

	// 인터페이스 타입의 변수이므로 implements한 구조체를 대입받을 수 있습니다.
	var starcraft Starcraft
	starcraft = &Protoss{"질럿"}
	starcraft.Attack() // 는 대입된 구조체의 메서드를 호출합니다.

	starcraft = &Terran{"마린"}
	starcraft.Attack()

	starcraft = &Zerg{"저글링"}
	starcraft.Attack()
}
