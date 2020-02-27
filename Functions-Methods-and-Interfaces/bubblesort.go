package main

import "fmt"

func Swap(a []int, i int){
	a[i], a[i+1] = a[i+1], a[i]
}

func BubbleSort(a []int){
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				Swap(a, j)
			}
		}
	}
}

func main() {

	fmt.Println("Please enter 10 integers: ")
	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}

	BubbleSort(a)
	fmt.Println(a)
}

