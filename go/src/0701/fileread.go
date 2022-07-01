package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var str string

	file1, _ := os.Open("test_1.txt")
	defer file1.Close()
	fmt.Fscanln(file1, &num1, &num2, &str)
	fmt.Println(num1, num2, str)
	file2, _ := os.Open("test_2.txt")
	defer file2.Close()
	fmt.Fscanln(file2, &num1, &num2, &str)
	fmt.Println(num1, num2, str)
}
