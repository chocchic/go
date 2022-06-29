package main

import . "fmt"

func main() {

	// 3개의 데이터를 저장할수 있는 slice 생성
	var li = make([]string, 3)
	list := []string{"참치"}
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

	// 그룹별 데이터 저장 - 배열(슬라이스)을 배열(슬라이스)로
	// 배열은 인덱스를 이용해서 저장하고 map은 key를 이용해서 저장합니다.
	// 세부 데이터를 배열을 이용해서 저장하는 것보다는 Map을 이용해서 세부 데이터를 저장하고 이들의 list를 이용하는 방법
	// MVC : 데이터를 생성하는 부분과 출력하는 부분을 분리하고 이를 Controller를 이용해서 연결하는 방식

	// 팀별 야구선수 명단을 저장

	// 데이터 생성
	haitai := []string{"선동렬", "이종범", "최향남", "이대진", "김상진"}
	kia := []string{"양현종", "이범호", "김주찬", "최형우"}

	//키는 문자열이고 데이터는 slice인 map
	kbo := map[string][]string{
		"해태": haitai,
		"기아": kia,
	}

	haitai = append(haitai, "김진우")
	kbo["빙그레"] = []string{"이강돈", "이정훈"}

	// map은 레퍼런스 타입이므로 참조를 대입하면 참조가 복사됩니다.
	baseball := kbo
	baseball["빙그레"] = []string{"이상군", "한희민"}
	// 데이터 출력
	for key, value := range kbo {
		Printf("%s\n", key)
		for _, player := range value {
			Printf("%s\t", player)
		}
		Println()
	}

}
