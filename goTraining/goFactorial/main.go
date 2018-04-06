package main

import (
	"runtime"
	"fmt"
)

func main() {
	for N := 1; N <= 13; N++ {
		fmt.Printf("%v = factorial(%v)\n", factorial(N), N)
		f := 1
		for i := 2; i <= N; i++ {
			f *= i
		}
		fmt.Printf("%v = %v!\n", f, N)
	}
}

func factorial(n int) int32 {
	m := runtime.NumCPU() - 1 // for main
	if (n < m) {
		m = n
	}
//	fmt.Println("m = ", m)
	term := make(chan int)
	done := make(chan bool)
	stride := n / m
//	fmt.Println("stride = ", stride)
	go func() {
		for i := 0; i < m; i++ {
			<- done
		}
		close(term)
	}()
	extra := n % m
//	fmt.Println("extra = ", extra)
	h := 0
	for i := 0; i < m; i++ {
		l := h + 1
//		fmt.Println("l = ", l)
		h = h + stride
		if extra > 0 {
			h++
			extra--
		}
//		fmt.Println("h = ", h)
		go func(low int, high int){
			var result int32 = 1
			for i := low; i <= high; i++ {
//				fmt.Printf("%v-%v: i = %v\n", low, high, i)
				result *= int32(i)
//				fmt.Println("result = ", result)
			}
//			fmt.Printf("%v-%v: result = %v\n", low, high, result)
			term <- int(result)
			done <- true
		}(l, h)
	}
	var result int32 = 1
	for x := range term {
//		fmt.Println("x = ", x)
		result *= int32(x)
	}
	return result
}