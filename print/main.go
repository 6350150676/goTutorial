package main

import "fmt"

func main() {
	age := 25
	name := "ramesh"
	height := 165.56

	fmt.Println(age, height, name)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("Type of name is %T\n", height)
	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
}
