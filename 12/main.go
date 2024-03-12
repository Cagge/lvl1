package main

import "fmt"

type set[T comparable] map[T]struct{}

func NewSet[T comparable](list []T) *set[T] {
	s := make(set[T])

	for _, val := range list {
		if _, ok := s[val]; !ok {
			s[val] = struct{}{}
		}
	}

	return &s
}

func main() {
	list := []string{"cat", "cat", "dog", "cat", "tree"}

	s := NewSet(list)

	fmt.Println(s)
}
