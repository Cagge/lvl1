package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // создаем канал

	// Постоянная запись данных в канал из главного потока
	go func() {
		for i := 1; ; i++ {
			ch <- fmt.Sprintf("Данные %d", i) // записываем произвольные данные в канал
			time.Sleep(1 * time.Second)       // ждем 1 секунду
		}
	}()

	numWorkers := 3 // количество воркеров
	for i := 0; i < numWorkers; i++ {
		// Набор из N воркеров, которые читают данные из канала и выводят их в stdout
		go func(workerID int) {
			for data := range ch {
				fmt.Printf("Воркер %d: %s\n", workerID, data) // выводим данные в stdout
			}
		}(i)
	}
	fmt.Scanln()
}
