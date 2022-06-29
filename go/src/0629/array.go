package main

import (
	"fmt"
	. "fmt"
)

func main() {
	// 3개의 데이터가 생성됨 - 기본적으로 초기화가 수행됨
	var ar [3]int
	// 참조를 저장할 준비만 됨
	var slice []int

	Println(ar)
	Println(slice)

	// 4개의 데이터를 저장할 수 있는 슬라이스를 생성해서 대입
	slice = make([]int, 4)
	Println(slice)

	a := []string{"딸기", "레몬", "수박", "참외"}
	Println(a[0])

	for _, v := range a {
		Println(v)
	}

	fmt.Println("====== 배열과 다른 점=====")
	xr := [3]string{"파파존스", "도미노", "피자헛"}
	list := []string{"네네치킨", "굽네치킨", "노랑통닭"}

	// 배열을 다른 배열에 대입 - 배열의 데이터가 복제가 되기 때문에 yr이 데이터를 수정해도 xr은 변경되지 않음
	// 메모리를 2배로 사용하지만 외부의 영향을 받지 않음
	yr := xr
	yr[2] = "피자스쿨"
	Println(xr)

	// slice를 다른 slice에 대입 - 복제가 되지 않고 참조를 넘기기 때문에 ar을 변경하면 list에도 변경된 내용이 적용
	// 메모리를 적게 사용하지만 외부의 영향을 받음 - 가독성이 떨어질 수 있음
	al := list
	al[2] = "푸라닭"
	Println(list)

	// 3개의 데이터를 저장할수 있는 slice 생성
	var li = make([]string, 3)

	// slice의 데이터를 li에 복제
	copy(li, list)
	li[0] = "BHC"
	Println(list)

	// slice에 데이터 추가
	list = append(list, "김밥천국")
	list = append(list, "등촌칼국수", "지코바")
	Println(list)

	// 슬라이스를 추가
	st := []string{"당근", "오이", "양파"}
	list = append(list, st...)
	Println(list)

}
