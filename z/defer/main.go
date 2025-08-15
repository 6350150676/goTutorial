package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Starting of the program")
	data := add(9, 9)
	defer fmt.Println(data)
	fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}
