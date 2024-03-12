package main

import "fmt"

func quicksort(list []int) {
	// если в слайсе 1 или меньше элементов, сортировать уже нечего
	if len(list) <= 1 {
		return
	}

	// задаем крайний левый и крайний правый индексы
	left, right := 0, len(list)-1

	// выбираем опорный элемент
	def := len(list) / 2

	// перемещаем опорный элемент в конец слайса
	list[def], list[right] = list[right], list[def]
	// не забываем свапнуть сами индексы
	def, right = right, def

	// итерируемся по слайсу
	// если i-ый элемент меньше опорного, то мы помещаем его на позицию left и увеличиваем значение left
	for i := 0; i < len(list); i++ {
		if list[i] < list[def] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}

	// перемещаем опорный элемент после наименьшего значения
	list[left], list[def] = list[def], list[left]
	// не забываем свапнуть сами индексы
	left, def = def, left

	// сортируем части слайса, находящиеся слева и справа от опортного элемента
	quicksort(list[:def])
	quicksort(list[def+1:])

}

func main() {
	list := []int{-4, 2, 12, -5, 0, 2, -100, 1, -5}
	fmt.Println(list)

	quicksort(list)
	fmt.Println(list)
}
