package main

import (
	"fmt"
	"runtime"
	"sync"
)

var Zähler int
var wg sync.WaitGroup

func hochzählen() {
	temp := Zähler
	runtime.Gosched()
	//for i := 1; i < 11; i++ {
	//fmt.Printf("A")
	//}
	//fmt.Println()
	temp++
	Zähler = temp
	wg.Done()
}

func main() {

	fmt.Println(Zähler)

	wg.Add(100)
	for j := 0; j < 100; j++ {
		go hochzählen()
	}
	wg.Wait()
	fmt.Println(Zähler)

}
