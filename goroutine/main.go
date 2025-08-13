package main

import (
	"fmt"
	"time"
)

func sayhello() {
	fmt.Println("hello")
	time.Sleep(2000 * time.Millisecond)
}
func sayhi() {
	fmt.Println("hanji")
}
func main() {
	fmt.Println("learning groutine")
	go sayhello()
	go sayhi()
	time.Sleep(3000 * time.Millisecond)
}
