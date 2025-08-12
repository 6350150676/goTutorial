package main

import "fmt"

func main() {
	numbes := []int{1, 2, 23}
	fmt.Println(numbes)
	fmt.Printf("Number has data type: %T\n", numbes)
	numbes = append(numbes, 3, 10, 12, 13, 14, 15, 16, 17, 18, 19)
	fmt.Println(numbes)
	stock := make([]int, 0)
	fmt.Println("Slice:", stock)
}
