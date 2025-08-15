package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Numbers is :", i)
	}
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index of Numbers is %d, and value is %d\n", index, value)
	}
}
