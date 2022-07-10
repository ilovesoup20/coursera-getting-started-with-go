package main

import "fmt"

func main() {
	var num float32
	fmt.Println("Enter a floating number:")
	fmt.Scan(&num)
	fmt.Printf("You entered: %.2f\n", num)
	fmt.Printf("As truncated: %d\n", int(num))
}
