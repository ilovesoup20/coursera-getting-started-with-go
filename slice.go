package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	var input string

	for input != "X" {
		fmt.Println("Enter a number:")
		fmt.Scan(&input)

		intValue, _ := strconv.Atoi(input)
		slice = append(slice, intValue)

		sort.Ints(slice)
		fmt.Printf("Input: %d, Slice: %v\n", intValue, slice)

	}
}
