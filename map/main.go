package main

import "fmt"

func main() {
	//Example 1
	// name <-> grade
	studentGrades := make(map[string]int)
	studentGrades["Amit"] = 90
	studentGrades["Ankit"] = 80
	studentGrades["Ram"] = 70
	studentGrades["Shyam"] = 60
	fmt.Println("Mark of Amit is", studentGrades["Amit"])
	fmt.Println("All the students are", studentGrades)

	studentGrades["Ram"] = 100
	fmt.Println("Mark of Ram is", studentGrades["Ram"])
	delete(studentGrades, "Ram")
	fmt.Println("Marks of Ram are", studentGrades["Ram"])

	// Checking if a key exists
	grades, exists := studentGrades["David"]

	fmt.Println("Mark of David is", grades)
	fmt.Println("Is David present?", exists)

	//fmt.Println("Marks of David : ", studentGrades["David"])

	marks, Exists := studentGrades["Amit"]

	fmt.Println("Mark of David is", marks)
	fmt.Println("Is David present?", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and value is %d", index, value)
	}

	person := map[string]int{
		"Salman":   95,
		"Shahrukh": 80,
		"Akshay":   70,
		"Ameer":    90,
	}

	for index, value := range person {
		fmt.Printf(" -------Key is %s and value is %d", index, value)
	}

}
