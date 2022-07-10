package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animal Animal
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("> ")
		if scanner.Scan() {
			input = scanner.Text()
		}
		split := strings.Split(input, " ")

		switch split[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		}

		switch split[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
	cow.Eat()
	bird.Eat()
	snake.Eat()
}
