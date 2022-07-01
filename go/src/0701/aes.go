package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	// 해시 값을 추출하고자 하는 데이터를 생성
	s := "Hello, world! 12"
	key := "Hello, Key 12345"

	block, err := aes.NewCipher([]byte(key))
	if err != nil { // cipher가 안만들어졌으면 이 조건문 이후 실행 x
		fmt.Println(err)
		return
	}

	// 암호화
	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	// 복호화
	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext)
	fmt.Println(string(plaintext))
}
