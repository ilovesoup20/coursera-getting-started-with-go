package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	var filename string
	// filename := "names.txt"

	fmt.Println("Enter file name:")
	fmt.Scanln(&filename)

	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	names := make([]Name, 0, 3)

	for scanner.Scan() {

		fullname := scanner.Text()
		split := strings.Split(fullname, " ")

		name := Name{fname: split[0], lname: split[1]}
		names = append(names, name)
	}

	file.Close()

	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}
}
