package main

import "fmt"

func main() {
	num := 2
	ptr := &num

	fmt.Println(num) // prints 2
	fmt.Println(ptr) // prints the memory address
}
