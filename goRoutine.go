package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(string(v))
	}
}

func PrintNumber() {
	for i := 1; i <= 10; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func GetNumber() int {
	time.Sleep(1000 * time.Millisecond)
	return 3
}

func Callback(number int) {
	fmt.Printf("Callback value = %d", number)
}

func main() {
	wg.Add(2)

	go PrintHangul()
	go PrintNumber()
	go Callback(GetNumber())

	wg.Wait()

}
