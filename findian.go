package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)
	fmt.Printf("You entered %s\n", input)
	// input = strings.ToLower(input)

	if strings.IndexRune(input, 'i') == 0 && strings.IndexRune(input, 'n') == len(input)-1 && strings.IndexRune(input, 'a') > -1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
