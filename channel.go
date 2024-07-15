package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // channel을 닫아주지 않으면 무한정 대기하게 됨 -> deadlock
		fmt.Printf("Square: %d * %d = %d\n", n, n*n, n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
