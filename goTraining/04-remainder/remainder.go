package main

import "fmt"

func main() {
	var dividend int;
	fmt.Print("\nEnter a Large Integer: ")
	fmt.Scan(&dividend)

	var divisor int;
	fmt.Print("Enter a Smaller Integer: ")
	fmt.Scan(&divisor)
	
	var remainder int = dividend % divisor
	fmt.Printf("Remainder = %v\n\n", remainder)
}
