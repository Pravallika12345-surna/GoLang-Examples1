package main

import "fmt"

func main() {
	m1 := new(int)
	n1 := new(int)
	fmt.Println(m1, n1)
	somefunc(m1, n1)

	a := new(int)
	fmt.Printf("%T", a)

}

func somefunc(aptr *int, bptr *int) {

}
