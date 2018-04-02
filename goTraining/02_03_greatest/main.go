package main

import (
	"fmt"
	"math"
)

func greatest(a ...int) (g int) {
	g = math.MinInt32
	fmt.Println(g)
	for _, v := range a {
		if v > g {
			g = v
		}
	}
	return
}

func main() {
	g := greatest(-1000, 1000, 40, -5, 2)
	fmt.Println(g)
}