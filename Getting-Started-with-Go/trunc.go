package main

import (
    "fmt"
    "os"
)

func main() {
    var input float32
    var trunc int

    fmt.Printf("Enter a floating number: ")
    num, err := fmt.Scan(&input)

    if err != nil || num != 1 {
	fmt.Printf("Your input cannot be converted: %s\n", err)
	os.Exit(1)
    }

    trunc = int(input)
    fmt.Printf("Your input as an integer: %d\n", trunc)
}
