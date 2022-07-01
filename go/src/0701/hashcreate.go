package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	// 해시 값을 추출하고자 하는 데이터를 생성
	s := "chocchic"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1)

	sha := sha512.New()
	sha.Write([]byte("chocchic"))
	sha.Write([]byte("촉촉한초코칩"))
	h2 := sha.Sum(nil)
	fmt.Printf("%x\n", h2)
}
