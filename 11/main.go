package main

import "fmt"

type set[T comparable] map[T]struct{}

func main() {
	set1 := set[int]{
		2: {},
		1: {},
		4: {},
		3: {},
	}

	set2 := set[int]{
		4: {},
		2: {},
		6: {},
		8: {},
	}

	intersect := make(set[int])
	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersect[k] = struct{}{}
		}
	}

	fmt.Println(intersect)
}
