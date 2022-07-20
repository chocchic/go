package main

import (
	"fmt"
)

func toByte(s string) []byte {
	result := make([]byte, 0, len(s)/2+1)
	for i := 0; i < len(s); i += 2 {
		var curbyte byte = 0
		x := s[i : i+2]
		for i, xi := range x {
			if 48 <= xi && xi <= 57 {
				temp := byte((xi - 48)) << ((1 - i) * 4)
				curbyte += temp
			} else if 65 <= xi && xi <= 90 { // x의 i번째가 소문자일 때
				temp := byte((xi - 55)) << ((1 - i) * 4)
				curbyte += temp
			} else { //if 97 <= xi && xi <= 122 // x의 i번째가 대문자일 때
				temp := byte((xi - 87)) << ((1 - i) * 4)
				curbyte += temp
			}
		}
		result = append(result, curbyte)
	}

	return result
}

func main() {
	ciphertext := toByte("f03963DefABcD42b55060a6f688025b7de0b31777614ce174b835651843e301b64a52212d3226adc23a4545f1b204358bda427530920")

	for i, v := range ciphertext {
		fmt.Printf("ciphertext[%d] = %#x\n", i, v)
	}

}
