package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct{
	firstName []byte
	lastName []byte
}

func main() {

	persons := make([]Person, 0, 0)

	var fileName string
	fmt.Println("Enter file name: ")
	fmt.Scanln(&fileName)

	file, _ := os.Open(fileName)
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		s := strings.Split(s.Text(), " ")
		p := Person{firstName: []byte(s[0]), lastName: []byte(s[1])}

		persons = append(persons, p)	
	
	}

	for _, p := range persons {
		fmt.Printf("First Name: %s, Last Name: %s\n", p.firstName, p.lastName)
	}

}
