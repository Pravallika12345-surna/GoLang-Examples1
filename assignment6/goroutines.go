package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU Processing : ", runtime.NumCPU())
	wg.Add(2)
	go helloworld()
	go helloworld1()
	fmt.Println("after go")
	wg.Wait()
}

func helloworld() {
	defer wg.Done() //this line will get executed after executing the entire method
	fmt.Println("from hello one!")
	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}
	fmt.Println("from hello two!")

}

func helloworld1() {
	defer wg.Done() //this line will get executed after executing the entire method
	fmt.Println("from hello11!")
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println("from hello112!")
}
