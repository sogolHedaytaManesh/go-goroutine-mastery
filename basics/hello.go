package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello from", name)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go greet("Goroutine 1")
	go greet("Goroutine 2")

	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
