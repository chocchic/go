package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 문자열을 키로 해서 모든 자료형의 데이터를 저장할 수 있는 Map을 생성
	data := make(map[string]interface{})

	data["name"] = "choc"
	data["age"] = "25"

	doc, _ := json.Marshal(data)
	fmt.Println(doc)
}
