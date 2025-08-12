package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web services")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error ka jahtka", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)
	data, error1 := io.ReadAll(res.Body)
	if error1 != nil {
		fmt.Println("error ka jahtka", error1)
		return
	}
	fmt.Println("response:", string(data))
}
