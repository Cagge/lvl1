package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	mapa := make(map[int]int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mapa[i] = i
			fmt.Println(mapa[i])
		}(i)
	}
	wg.Wait()
	fmt.Println(mapa)
}
