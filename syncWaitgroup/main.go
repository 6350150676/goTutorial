package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("woerker %d started\n", i)
	fmt.Printf("worker %d end\n", i)
}
func main() {
	fmt.Println("wait")
	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i := 0; i <= 3; i++ {
		wg.Add(1) //increment the waitgroup counter
		go worker(i, &wg)
	}
	//wait for all the workers to finish 
	wg.Wait()
	fmt.Println("worker task completed ")
}
