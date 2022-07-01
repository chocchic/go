package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `{
		"name" : "chocchoc",
		"age" : 25,
		"color" : "white"
	}`

	// 파싱한 내용을 저장하기 위한 map을 생성
	var data map[string]interface{}

	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data["name"], data["age"], data["color"])
}
