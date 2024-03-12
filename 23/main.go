package main

import (
	"errors"
	"fmt"
)

func delete(list *[]int, idx int) error {
	l := *list
	if idx < 0 || idx >= len(l) {
		return errors.New("index out of range")
	}
	*list = append(l[:idx], l[idx+1:]...)
	return nil
}

func main() {
	list := []int{-4, 2, 12, -5, 0, 2, -100, 1, -5}
	target := 4
	err := delete(&list, target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
