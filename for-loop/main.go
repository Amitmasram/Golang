package main

import (
	"fmt"
)

func main() {

	// Example 1 : Basic For Loop
	for i := 0; i < 10; i++ {
		fmt.Println("Numbers is :", i)
	}

	// Example 2 : For Loop with break statement
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 5 {
			break
		}
	}

	// Example 3 : For Loop with Range Loop
	numbers := []int{13, 23, 3, 43, 5, 63}

	for index, value := range numbers {
		fmt.Printf("Index is %d Value is %d\n", index, value)
		fmt.Println("Index is", index, "Value is", value)
	}

	data := "Hello World"

	for index, char := range data {
		fmt.Printf("Index of data is %d Value is %c\n", index, char)
	}

	// Example 4 : For Loop with continue statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("Numbers is :", i)
	}
}
