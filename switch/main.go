package main

import "fmt"

func main() {

	// Example 1 : Basic Switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("No such day")
	}

	// Example 2 : Switch with multiple cases
	month := "January"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter Season")

	case "April", "May", "June":
		fmt.Println("Spring Season")
	case "July", "August", "September":
		fmt.Println("Summer Season")
	default:
		fmt.Println("Other Season")

	}

	// Example 3 : Switch with no expression
	var name string
	switch {
	case name == "Amit":
		fmt.Println("Hello Amit")
	case name == "Ankit":
		fmt.Println("Hello Ankit")
	default:
		fmt.Println("Hello World")
	}
	temperature := 30
	switch {
	case temperature < 0:
		fmt.Println("It's Freezing")
	case temperature >= 0 && temperature <= 10:
		fmt.Println("It's cold")
	case temperature > 10 && temperature <= 20:
		fmt.Println("It's normal")
	case temperature > 20 && temperature <= 30:
		fmt.Println("It's Warm")
	default:
		fmt.Println("It's too hot")

	}

	// Example 4 : Switch with fallthrough
	switch {
	case name == "Amit":
		fmt.Println("Hello Amit")
		fallthrough
	case name == "Ankit":
		fmt.Println("Hello Ankit")
		fallthrough
	default:
		fmt.Println("Hello World")
	}

	// Example 5 : Switch with true or false
	switch false {
	case name == "Amit":
		fmt.Println("Hello Amit")
	case name == "Ankit":
		fmt.Println("Hello Ankit")
	default:
		fmt.Println("Hello World")
	}

	// Example 6 : Switch with multiple expressions
	switch name, age := "Amit", 25; {
	case name == "Amit" && age == 25:
		fmt.Println("Hello Amit")
	case name == "Ankit" && age == 25:
		fmt.Println("Hello Ankit")
	default:
		fmt.Println("Hello World")
	}

}
