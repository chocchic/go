package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("test_1.txt")
	defer file1.Close()
	fmt.Fprintln(file1, 1, 1.7, "촠촉한 초코칩")

	file2, _ := os.Create("test_2.txt")
	defer file2.Close()
	fmt.Fprintf(file2, "%d %f %s", 2, 3.141592, "Pie")
}
