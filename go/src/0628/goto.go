package main

import . "fmt"

func main() {
	Println("프로그램1")
	goto ERROR
	Println("여긴 안 찍힘")
ERROR:
	Println("이곳으로 이동")
}
