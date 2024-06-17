package main

import (
	"fmt"
	"time"
)

func main() {
	// "dd-mm-yyyy"
	currentTime := time.Now()

	fmt.Println("Current time is", currentTime)
	fmt.Printf("Type of currentTime is %T\n", currentTime)

	formatted := currentTime.Format("02/01/2006,3:04 PM")

	fmt.Println("Formatted date is", formatted)
	layout_str := "2006-01-02"
	dateStr := "2024-10-17"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted date is", formatted_time)

	// Add one more day in current time
	new_date := currentTime.Add(48 * time.Hour)
	fmt.Println("Current time is", new_date)

	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("Formatted date is", formatted_new_date)

}
