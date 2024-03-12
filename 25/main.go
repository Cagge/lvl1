package main

import (
	"fmt"
	"time"
)

func sleep(sec int) {

	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	now := time.Now()
	sleep(4)
	fmt.Println(time.Since(now))
}
