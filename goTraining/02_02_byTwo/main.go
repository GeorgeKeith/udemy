package main

import "fmt"

func main() {
	byTwo := func (x int) (d int, isEven bool) {
		d = x / 2
		isEven = (x % 2) == 0
		return
	}
	d, isEven := byTwo(99)
	fmt.Println(d, isEven)

	d, isEven = byTwo(200)
	fmt.Println(d, isEven)
}
