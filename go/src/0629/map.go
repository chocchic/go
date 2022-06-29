package main

import . "fmt"

func main() {
	dic := map[string]string{"바나나": "banana", "사과": "apple", "당근": "carrot", "개": "dog"}
	// 하나의 값 출력
	Println(dic["바나나"])
	Println(dic["돌고래"]) // 없어서 아무 것도 안나옴

	// value에는 데이터가 존재하면 데이터, 데이터가 존재하지 않으면 기본값이 대입
	// ok에는 데이터의 존재 여부가 bool로 대입
	value, ok := dic["당근"]
	if ok == true {
		Println(value)
	} else {
		Println("데이터가 존재하지 않습니다.")
	}

	// 데이터 추가 및 업데이트
	dic["달걀"] = "egg"
	Println(dic)

	// 데이터 삭제
	delete(dic, "개")
	Println(dic)

	// 데이터 전체 순회
	for key, value := range dic {
		Printf("%s : %s\n", key, value)
	}
}
