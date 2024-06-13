package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	return a + b
}
func multiplication(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("Hello, World!")

	simpleFunction()
	ans := add(1, 2)
	fmt.Println(ans)

	ans = multiplication(1, 2)
	fmt.Println(ans)
}
