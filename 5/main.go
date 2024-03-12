package main

import (
	"fmt"
	"os"
	"time"
)

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func receiver(ch chan int) {
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, exiting receiver")
			return
		}
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int)

	go sender(ch)
	go receiver(ch)

	time.Sleep(5 * time.Second) // Задержка в 5 секунд
	os.Exit(0)
	fmt.Println("Main goroutine exiting")
}
