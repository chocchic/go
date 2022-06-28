package main

import (
	. "fmt"
)

func main() {
	const max int = 10
	//max = 20
	Println(max)

	const (
		MIN    = 0
		NORMAL = 5
		MAX    = 10
	)
	Println(MAX)

	const (
		ZERO = iota
		ONE
		TWO
	)
	Println(TWO)

	const (
		LBUTTON = 1 << iota
		RBUTTON = 1 << iota
		SHIFT   = 1 << iota
		CONTROL = 1 << iota
	)
	Println(CONTROL)

	const (
		SALES    = iota * 10
		ACCOUNT  = iota * 10
		RESEARCH = iota * 10
	)
	Println(RESEARCH)
}
