package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter 10 numbers:")

	var input string
	// input = "5 4 3 9 8 1 2 7 6 0"

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	numbersSplit := strings.Split(input, " ")
	numbers := make([]int, 0, 10)

	for i := 0; i < len(numbersSplit); i++ {
		n, _ := strconv.Atoi(numbersSplit[i])
		numbers = append(numbers, n)
	}

	BubbleSort(numbers)
	fmt.Println(numbers)
}

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for y := 0; y < len(numbers)-1; y++ {
			curr := numbers[y]
			next := numbers[y+1]
			if curr > next {
				Swap(numbers, y)
			}
		}
	}
}

func Swap(slice []int, index int) {
	tmp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = tmp
}
