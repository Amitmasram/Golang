package main

import "fmt"

func simpleFunction(){
	fmt.Println("simple function")
}

func add(a,b int)(int){
	return a+b
}
func main() {
	fmt.Println("Hello, World!")

	simpleFunction()
	ans := add(1,2)
	fmt.Println(ans)
}