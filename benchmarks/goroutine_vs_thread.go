package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func goroutineTask(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	num := 100000

	start := time.Now()
	for i := 0; i < num; i++ {
		wg.Add(1)
		go goroutineTask(&wg)
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Created %d goroutines in %s\n", num, elapsed)
}
