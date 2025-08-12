package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "jahtke"
	name[2] = "maar"
	fmt.Println(name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is:", numbers)

	var price [5]int
	fmt.Println("Price is :", price)
}
