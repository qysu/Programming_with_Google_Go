package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal interface{
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c *Cow) Eat(){fmt.Println("grass")}
func (c *Cow) Move(){fmt.Println("walk")}
func (c *Cow) Speak(){fmt.Println("moo")}

func (b *Bird) Eat(){fmt.Println("worms")}
func (b *Bird) Move(){fmt.Println("fly")}
func (b *Bird) Speak(){fmt.Println("peep")}

func (b *Snake) Eat(){fmt.Println("mice")}
func (b *Snake) Move(){fmt.Println("slither")}
func (b *Snake) Speak(){fmt.Println("hsss")}

func main(){
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSpace(input), " ")

		var a Animal
		switch s[0] {
		case "cow":
			a = new(Cow)
		case "bird":
			a = new(Bird)
		case "snake":
			a = new(Snake)
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
