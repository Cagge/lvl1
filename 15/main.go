package main

var justString string

// Проблема 1: срез строк сломается, если символы выйдут за 1 бит.
// Проблема 2: глобальная переменная justString висит в памяти, хотя мы используем только первые 100 элементов.
// Решение: создание slice рун из строки v
func someFunc1() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func someFunc2() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc1()
	someFunc2()
}
