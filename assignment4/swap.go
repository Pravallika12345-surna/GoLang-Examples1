package main

import "fmt"

func swapnum(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x, y := 5, 10
	fmt.Println("Before SWAP: ", x, y)
	swapnum(&x, &y)
	fmt.Println("After SWAP:  ", x, y)
}
