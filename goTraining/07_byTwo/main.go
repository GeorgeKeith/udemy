package main

import "fmt"

func byTwo(x int) (d int, isEven bool) {
	d = x/2
	isEven = (x % 2) == 0
	return
}

func main() {
	d, isEven := byTwo(9)
	fmt.Println(d, isEven)
	d, isEven = byTwo(20)
	fmt.Println(d, isEven)
}