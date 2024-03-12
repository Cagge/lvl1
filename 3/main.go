package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	massiv := [5]int{2, 4, 6, 8, 10}
	var outlist [5]int
	for i, v := range massiv {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			outlist[i] = (v * v)
		}(v)
	}
	wg.Wait()
	fmt.Println(outlist)
	outint := 0
	for _, v := range outlist {
		outint += v
	}
	fmt.Println(outint)
}
