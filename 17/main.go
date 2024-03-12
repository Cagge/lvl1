package main

import "fmt"

func binarySearch(list []int, val int) int {
	// задаем крайний левый и крайний правый индексы
	left := 0
	right := len(list) - 1

	for left <= right {
		// находим индекс среднего элемента
		mid := (left + right) >> 1

		// если элемент по индексу mid равен искомому значению, возвращаем его
		if list[mid] == val {
			return mid
			// если он больше, сдвигаем правый индекс, чтобы он стал равен mid - 1
		} else if list[mid] > val {
			right = mid - 1
			// если он меньше, сдвигаем левый индекс, чтобы он стал равен mid + 1
		} else {
			left = mid + 1
		}
	}
	// если искомого элемента нет в массиве - возвращаем -1
	return -1
}

func main() {
	list := []int{-4, 2, 12, -5, 0, 2, -100, 1, -5}
	val := -2
	fmt.Println(binarySearch(list, val))
}
