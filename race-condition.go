package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines before start:", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Printf(" C: %v ", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Printf("R: %v ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("\nGoroutines after running:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
