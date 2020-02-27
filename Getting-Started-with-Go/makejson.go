package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	m := make(map[string] string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name: ")
	name, err := reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Enter address: ")
	address, err := reader.ReadString('\n')

        if err != nil {
                os.Exit(1)
        }

	m["name"] = strings.Trim(name, " \n")
	m["address"] = strings.Trim(address, " \n")

	jsonString, err := json.MarshalIndent(m, "", "\t")
	
	if err != nil {
                os.Exit(1)
        }

	fmt.Printf("\n%s\n", jsonString)
}
