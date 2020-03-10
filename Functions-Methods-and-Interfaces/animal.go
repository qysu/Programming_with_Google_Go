package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct{
	food string
	locomotion string
	sound string
}

func (a Animal) Eat(){
	fmt.Println(a.food)
}

func (a Animal) Move(){
	fmt.Println(a.locomotion)
}

func (a Animal) Speak(){
	fmt.Println(a.sound)
}

func main(){
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSpace(input), " ")

		var a Animal
		switch s[0] {
		case "cow":
			a.food = "grass"
			a.locomotion = "walk"
			a.sound = "moo"
		case "bird":
			a.food = "worms"
			a.locomotion = "fly"
			a.sound = "poop"
		case "snake":
			a.food = "mice"
			a.locomotion = "slither"
			a.sound = "hsss"
		}

		switch s[1] {
		case "eat":
			a.Eat()
		case "speak":
			a.Speak()
		case "move":
			a.Move()
		}
	}
}