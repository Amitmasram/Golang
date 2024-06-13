package main

import "fmt"

func main() {
	// Println
	fmt.Println("Hello, World!")

	// Printf
	name := "Alice"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of age is %T\n", age)
	
}