package main

import (
	"fmt"
)

func typeof(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan bool, chan any:
		return "chan"
	default:
		return "invalid type"
	}
}

func main() {
	var val interface{} = 1
	fmt.Println(typeof(val))

	val = "abc"
	fmt.Println(typeof(val))

	val = true
	fmt.Println(typeof(val))

	val = make(chan int)
	fmt.Println(typeof(val))
}
