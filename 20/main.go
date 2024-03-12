package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	// разбиваем строку слова

	list := strings.Split(str, " ")
	// создаем два маркера
	left, right := 0, len(list)-1

	// пока левый индекс меньше правого, меняем элементы массива местами
	for left < right {
		list[left], list[right] = list[right], list[left]

		// сдвигаем правый указатель на 1 влево, левый - на 1 вправо
		right--
		left++
	}

	// соединяем массив обратно в строку через пробел
	return strings.Join(list, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
