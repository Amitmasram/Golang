package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func divide2(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "cannot divide by zero"
	}
	return a / b, ""
}

func main() {
	fmt.Println("started Error Handling in the function")
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error Handling")

	// }
	// fmt.Println("Division of two numbers", ans)

	ans, _ := divide2(10, 0)
	fmt.Println("Division of two numbers", ans)

}
