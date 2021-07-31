package main

import "fmt"

func main() {
	var a *int
	b := 5
	a = &b
	fmt.Println(b)
	fmt.Printf("%v\n", a)

	var a1 *string
	b1 := "Hello"
	a1 = &b1
	fmt.Println(b1)
	fmt.Printf("%v\n", a1)

	var a2 *float64
	b2 := 10.0
	a2 = &b2
	fmt.Println(b2)
	fmt.Printf("%v", a2)
}
