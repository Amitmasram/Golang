package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple, orrange, banana, guava"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two five six six two"
	count := strings.Count(str, "two")
	fmt.Println(count)

	// Find accurance
	index := strings.Index(str, "two")
	fmt.Println(index)

	str = "                    Hello ,       Go       !             "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "Amit"
	str2 := "Masram"

	result := strings.Join([]string{str1, " S", str2}, " ")

	fmt.Println(result)
}
