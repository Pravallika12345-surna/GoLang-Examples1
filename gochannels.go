package main

import "fmt"

func main() {
	msg := make(chan string)
	go chan1(msg)
	fmt.Println(<-msg)
	msg1 := make(chan float64)
	go chan2(msg1)
	fmt.Println(<-msg1)
	msg2 := make(chan string)
	go chan3(msg2)
	fmt.Println(<-msg2)
	msg3 := make(chan int)
	go chan4(msg3)
	fmt.Println(<-msg3)

}

func chan1(msg chan string) {

	msg <- "Hello I am Form Channel1"
}
func chan2(msg1 chan float64) {
	msg1 <- 2.0
}
func chan3(msg2 chan string) {
	msg2 <- "I am from Channel3"
}

func chan4(msg3 chan int) {
	for i := 0; i < 4; i++ {
		msg3 <- i
	}

}
