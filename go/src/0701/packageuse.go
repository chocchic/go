package main

import (
	"calc" // 환경변수 설정에 gopath로 되어있는 디렉토리 내의 src파일에 calc라는 패키지를 저장해두어야 가능
	"fmt"
)

func main() {
	fmt.Println(calc.AddWithInteger(100, 200))
}
