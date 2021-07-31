package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendvalue(c chan int) {
	defer wg.Done()
	fmt.Print("\nBuffer Channel :")
	for i := 0; i <= 10; i++ {
		fmt.Print(" ", i)
		c <- i
	}
}

func main() {

	msgValue := make(chan int, 10)
	defer close(msgValue)

	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)
	go sendvalue(msgValue)

	values := msgValue
	fmt.Println(values)

}
