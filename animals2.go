package animals2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalInfo struct {
	food       string
	locomotion string
	noise      string
}

// Cow
type Cow AnimalInfo

func (cow *Cow) Eat() {
	fmt.Println(cow.food)
}
func (cow *Cow) Move() {
	fmt.Println(cow.locomotion)
}
func (cow *Cow) Speak() {
	fmt.Println(cow.noise)
}

// Bird
type Bird AnimalInfo

func (bird *Bird) Eat() {
	fmt.Println(bird.food)
}
func (bird *Bird) Move() {
	fmt.Println(bird.locomotion)
}
func (bird *Bird) Speak() {
	fmt.Println(bird.noise)
}

// Snake
type Snake AnimalInfo

func (snake *Snake) Eat() {
	fmt.Println(snake.food)
}
func (snake *Snake) Move() {
	fmt.Println(snake.locomotion)
}
func (snake *Snake) Speak() {
	fmt.Println(snake.noise)
}

func main() {
	animals := make(map[string]Animal)

	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("> ")
		if scanner.Scan() {
			input = scanner.Text()
		}

		tokens := strings.Split(input, " ")

		command := tokens[0]
		name := tokens[1]

		if command == "newanimal" {
			animalType := tokens[2]
			var newanimal Animal

			switch animalType {
			case "cow":
				newanimal = &Cow{"grass", "walk", "moo"}
			case "bird":
				newanimal = &Bird{"worms", "fly", "peep"}
			case "snake":
				newanimal = &Snake{"mice", "slither", "hsss"}
			default:
				fmt.Println(animalType, "is invalid")
				continue
			}

			animals[name] = newanimal

			fmt.Println("Created!")
		} else if command == "query" {
			action := tokens[2]

			animal := animals[name]

			if animal == nil {
				fmt.Println(name, "animal not found")
				continue
			}

			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid command", action)
			}
		}
	}
}
