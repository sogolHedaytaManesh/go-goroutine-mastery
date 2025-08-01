package main

import "fmt"

func deepRecursion(i int) {
	if i > 0 {
		deepRecursion(i - 1)
	} else {
		fmt.Println("Reached base case")
	}
}

func main() {
	go deepRecursion(1000)
	select {}
}
