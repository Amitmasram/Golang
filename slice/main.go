package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers = append(numbers, 8, 11, 13, 14, 15)

	fmt.Println(numbers)
	fmt.Printf("Type of numbers is %T\n", numbers)
	fmt.Println("Length of numbers is", len(numbers))
	name := []string{}
	fmt.Println(name)

	/// Make function
	num := make([]int, 3, 5)
	num = append(num, 4)
	num = append(num, 69)
	num = append(num, 6)

	fmt.Println("Slice is", num)
	fmt.Println("Length of slice is", len(num))
	fmt.Println("Capacity of slice is", cap(num))

	stock := make([]int, 0)
	fmt.Println("Slice is", stock)
	fmt.Println("Length of slice is", len(stock))
	fmt.Println("Capacity of slice is", cap(stock))

}
