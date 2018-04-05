package main

import "fmt"

func main() {
	var a int64 = 1
	var b int64 = 1
	var s int64 = 0
	for b < 4000000 {
		if (b % 2) == 0 {
			s += b
		}
		c := a + b
		a = b
		b = c
	}
	fmt.Println(s)
}
// Even Fibonacci numbers
// Problem 2