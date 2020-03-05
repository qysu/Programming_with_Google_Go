package main

import "fmt"

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		fmt.Println( a)
		fmt.Println( v)
		fmt.Println( s)
		return 0.5*a*t*t+v*t+s
	}
}

func main() {
	var a float64
	var v0 float64
	var s0 float64
	var t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fmt.Printf("Total displacement is %v. \n", fn(t))
}


