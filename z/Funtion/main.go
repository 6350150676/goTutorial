package main

import "fmt"

func sf() {
	println("gg")
}
func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("we are learning function in Golang")
	sf()
	add(1, 2)
}
