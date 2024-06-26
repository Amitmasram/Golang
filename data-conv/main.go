package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println(num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	data = data + 0.5
	fmt.Println("Data is ", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "123"
	number_int, _ := strconv.Atoi(number_string)
	number_int = number_int + 13533
	fmt.Println("number_int is ", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	num_string := "3.14"
	num_float, _ := strconv.ParseFloat(num_string, 64)
	num_float = num_float + 0.5
	fmt.Println("num_float is ", num_float)
	fmt.Printf("Type of num_float is %T\n", num_float)
}
