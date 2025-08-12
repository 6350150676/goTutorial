package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	// Example usage
	p := Person{Name: "Alice", Age: 25, IsAdult: true}
	fmt.Println(p)

	// Convert person into JSON (Marshalling)
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println(string(jsonData))

	// Convert JSON back into Person (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	fmt.Println(personData)
}
