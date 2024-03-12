package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var now time.Time = time.Now()
	// var year int = now.Year()
	// fmt.Println(year)
	// broken := "G# r#cks!"
	// replacer := strings.NewReplacer("#", "o")
	// fixed := replacer.Replace(broken)
	// fmt.Println(fixed)
	fmt.Println(UnikRun(5))
}

func UnikRun(find int) map[int]int {
	slice := make(map[int]int)

	for i := 0; i < find; i++ {
		temp := rand.Intn(5)
		for _, v := range slice {
			for {
				if v != temp {
					break
				} else {
					temp = rand.Intn(5)
				}
			}
		}
		slice[i] = temp

		// for i2 := 0; i2 < 100000; i2++ {
		// 	temp = rand.Intn(100)
		// 	if slice[i] != temp {
		// 		break
		// 	} else {
		// 		continue
		// 	}
		// }
		// slice = append(slice, temp)
	}
	return slice
}
