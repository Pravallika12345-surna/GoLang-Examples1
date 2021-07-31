package main

import "fmt"

func main() {

	num1 := []int{1, 2, 3, 4, 5, 6}
	num2 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(num1, reverse(num2))
	fmt.Println(add(num1, num2))
}

func reverse(a []int) []int {

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func add(a []int, num1 []int) []int {

	for i := 0; i < len(a); i++ {
		a[i] += num1[i]
	}
	return a

}

/*
func add1(a []int, num1 []int) string {
	c := make([]string, 10)
	for i := 0; i < len(a); i++ {
		//	c = string(a[i] + num1[i])
		c[i] = string(a[i] + num1[i])
	}
	return "c"

}  */
