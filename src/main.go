
package main

import (
	"fmt"
	
)
func main(){
	fmt.Println("Hello Amit How are you ?")
	var name string = "Amit";
	var version = "1.0.0";
	fmt.Println(name);
	fmt.Println(version);

	var money int = 50000;
	fmt.Println(money);
	var dimension float32 = 5.5;
    fmt.Println(dimension);
	var isOk bool = true;
	fmt.Println(isOk);

	const pi = 3.14;
	fmt.Println(pi);

	amit := "Amit";
	fmt.Println(amit);

	var (
		name1 string = "Amit";
		version1 = "1.0.0";
	)
	fmt.Println(name1);
	fmt.Println(version1);
	var Public = "data is ipublic";
	var Private = "data is private";
	fmt.Println(Public);
	fmt.Println(Private);
}