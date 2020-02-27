package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	// "strings"
)

func main() {
	slice_1 := make([]int, 0, 3)

	for {
		var input string

		fmt.Println("Enter number: ")
		fmt.Scan(&input)

		if input == "x" || input == "X" {
			break 
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			os.Exit(1)
		}

		slice_1 = append(slice_1, num)

		sort.Slice(slice_1, func(i, j int) bool {
			return slice_1[i] < slice_1[j]
		})
		
		fmt.Println(slice_1)

	}

}
