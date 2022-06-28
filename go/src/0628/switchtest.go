package main

import . "fmt"

func main() {
	value := 2
	switch value {
	case 1:
		Println("중식")
	case 2:
		Println("양식")
		fallthrough
	case 3:
		Println("한식")
	}
}
