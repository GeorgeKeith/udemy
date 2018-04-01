package main

import "fmt"

func main() {
	var yourName string

	fmt.Print("Enter your name: ")
	fmt.Scan(&yourName)
	fmt.Println("\nHello, " + yourName)
}
