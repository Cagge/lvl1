package main

import (
	"fmt"
	"strings"
)

func main() {
	testCase := []string{"abcd", "abCdefAaf", "aabcd"}
	for i, _ := range testCase {
		temp := strings.ToLower(testCase[i])
		_, out := CountLetters(temp)
		fmt.Println(testCase[i], out)
	}
}
func CountLetters(str string) (string, bool) {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return str, false
			}
		}
	}

	return str, true
}
