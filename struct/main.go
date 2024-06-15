package main

import "fmt"

// Define a struct name Person
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
	HouseNo int
	Area    string
	State   string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var amit Person
	amit.FirstName = "Amit"
	amit.LastName = "Kumar"
	amit.Age = 25
	fmt.Println("Amit Person :", amit)

	// 2nd method
	person1 := Person{
		FirstName: "Sumit",
		LastName:  "Masram",
		Age:       15,
	}

	fmt.Println("Person1 :", person1)

	// new Keyword
	var person2 = new(Person)
	person2.FirstName = "Dishita"
	person2.LastName = "Merkhed"
	person2.Age = 21
	fmt.Println("Person2 :", person2.Age)

	var employee1 Employee

	employee1.Person_Details = Person{
		FirstName: "Amit",
		LastName:  "Kumar",
		Age:       25,
	}

	employee1.Person_Contact.Email = "KsNjv@example.com"
	employee1.Person_Contact.Phone = "1234567890"

	employee1.Person_Address = Address{
		HouseNo: 12,
		Area:    "Gorewada, Pune",
		State:   "Maharashtra",
	}

	fmt.Println("Employee1 :", employee1)

}
