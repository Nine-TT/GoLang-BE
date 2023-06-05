package main

import (
	"fmt"
	"sync"
)

type Num struct {
	Num int
}

func Gr1(wg *sync.WaitGroup) {
	fmt.Println("Goroutine 1")
	wg.Done()

}

func Gr2(wg *sync.WaitGroup) {
	fmt.Println("Goroutine 2")
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go Gr1(&wg)
	go Gr2(&wg)

	wg.Wait()

}
