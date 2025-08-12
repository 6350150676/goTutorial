package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 0
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of Bob: ", studentGrades["Bob"])

	delete(studentGrades, "Bob")

	Grades, Exists := studentGrades["David"]
	fmt.Println("Grades of David: ", Grades)
	fmt.Println("Davis exists : ", Exists)

	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}
	for index, value := range person {
		fmt.Printf("key is %s and marks is %d\n", index, value)
	}
}
