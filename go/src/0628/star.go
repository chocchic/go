package main

import . "fmt"

func main() {
	Println("문제 1")
	for i := 0; i < 6; i++ {
		for j := 0; j < i; j++ {
			Print("*")
		}
		Print("\n")
	}
	for i := 0; i < 25; i++ {
		Print("*")
		if i%5 == 4 {
			Print("\n")
		}
	}
	Println("문제 2")
	for i := 5; i > 0; i-- {
		for j := i; j > 0; j-- {
			Print("*")
		}
		Print("\n")
	}
	Println("문제 3")
	for i := 1; i < 6; i++ {
		if i > 6/2 {
			for j := 6 - i; j > 0; j-- {
				Print("*")
			}
		} else {
			for j := i; j > 0; j-- {
				Print("*")
			}
		}
		Print("\n")
	}
	Println("문제 4")
	for i := 4; i >= 0; i-- {
		for k := i; k > 0; k-- {
			Print(" ")
		}
		for j := (5-i)*2 - 1; j > 0; j-- {
			Print("*")
		}
		Print("\n")
	}

	Println("문제 5")
	for i := 4; i >= 0; i-- {
		for j := i; j > 0; j-- {
			Print(" ")
		}
		Print("*")
		for k := (4-i)*2 - 1; k > 0; k-- {
			if i != 0 {
				Print(" ")
			} else {
				Print("*")
			}
		}
		if 4-i > 0 {
			Print("*")
		}
		Print("\n")
	}
}
