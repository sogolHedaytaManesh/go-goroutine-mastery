package main

import (
	"fmt"
	"time"
)

func increment(counter *int) {
	*counter++
}

func main() {
	counter := 0

	for i := 0; i < 1000; i++ {
		go increment(&counter)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Final Counter Value:", counter)
}
