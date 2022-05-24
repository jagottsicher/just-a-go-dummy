
package main

import "fmt"

func main() {
	kanalVonInts := make(chan int, 2)

	go func() {
		kanalVonInts <- 23
		kanalVonInts <- 24
		kanalVonInts <- 25
		kanalVonInts <- 26

	}()

	fmt.Println(<-kanalVonInts)
	fmt.Println(<-kanalVonInts)
	fmt.Println(<-kanalVonInts)
fmt.Println(<-kanalVonInts)

}
