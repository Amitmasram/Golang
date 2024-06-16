package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 2
}

func main() {
	// Decalare a variable
	var num int
	num = 10

	var ptr *int
	ptr = &num
	fmt.Println("Value of num is", num)
	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Value of ptr is", *ptr)

	// IN Short way

	num1 := 6

	ptr = &num1

	fmt.Println("Value of num1 is", num1)
	fmt.Println("Pointer contains", ptr)
	fmt.Println("Data contains through pointer", *ptr)

	var pointer *int // default pointer is nil
	if pointer == nil {
		fmt.Println("Pointer is nil")
	}

	value := 10
	modifyValueByReference(&value)

	fmt.Println("Value contains :", value)

	// Using fmt.Println and fmt.Printf
	fmt.Println("Printing a message using fmt.Println")
	fmt.Printf("Printing a message using fmt.Printf\n")
	fmt.Printf("The value of num is %d\n", num)
	fmt.Printf("The value of num1 is %d\n", num1)
	fmt.Printf("The value of value is %d\n", value)
}
