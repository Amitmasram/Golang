package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, World!")
	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Names of the students are:", arr)
	var name = [5]string{"Ram", "Shyam", "Hari", "Gita", "Sita"}

	name[0] = "Amit"
	name[1] = "Ankit"
	fmt.Println("Names of the students are:", name)
	fmt.Println("Length of the array is:", name[2])

	var price [5]int
	fmt.Println("Price is :", price)

	var price2 [5]bool
	fmt.Println("Price is :", price2)

	var price3 [10]string

	fmt.Println("Price is :", price3)
	fmt.Printf("Type of price is %q\n", price3)

	var students [5]string
	students[2] = "Ankit"
	students[0] = "Akash"
	fmt.Println("Names of the students are:", students)
	fmt.Printf("name of the students are %q\n", students)
}
