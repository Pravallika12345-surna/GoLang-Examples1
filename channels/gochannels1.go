package main

import (
	"fmt"
)

//var sy sync.WaitGroup

func main() {
	ch := make(chan int, 4)
	//sy.Add(2)
	go write(ch)
	fmt.Println(<-ch)
	go write1(ch)
	fmt.Println(<-ch)
	/*	go write2(ch)
		fmt.Print(<-ch) */
}

func write(ch chan int) {
	//defer sy.Done()
	fmt.Println("\n from Channel one: ")
	for i := 0; i <= 4; i++ {
		ch <- i
		fmt.Println(i)
	}
	//defer close(ch)
}
func write1(ch1 chan int) {
	//defer sy.Done()
	fmt.Print("\n from Channel two: ")
	for i := 0; i < 5; i++ {
		ch1 <- i
		fmt.Print(i)
	}
	//defer close(ch1)
} /*
func write2(ch chan int) {
	defer sy.Done()
	fmt.Println("\n from Channel three: ")
	for j := 0; j < 5; j++ {
		ch <- j
		fmt.Println(j)
	}

} */
