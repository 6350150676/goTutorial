package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House int
	Area  string
	State string
}
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var lav Person
	fmt.Println(lav)
	lav.FirstName = "Prince"
	lav.LastName = "Agarwal"
	lav.Age = 24
	fmt.Println(lav)

	personi := Person{
		FirstName: "Aakash",
		LastName:  "Singh",
		Age:       26,
	}
	fmt.Println(personi)
	// new keyword method
	var person2 = new(Person) //new kwyword gives pointer
	person2.FirstName = "Simarn"
	person2.LastName = "Agarwal"
	person2.Age = 26

	fmt.Println(person2)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Prince",
		LastName:  "Agarwal",
		Age:       26,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "98765432"
	employee1.Person_Address = Address{
		House: 12,
		Area:  "Ranchi",
		State: "Jharkhand",
	}

	fmt.Println("Employee 1 : ", employee1)
}
