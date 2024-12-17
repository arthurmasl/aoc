package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("src/2024/03/example")

	pattern := []byte("mul(D,D)")
	temp := make([]byte, 12)

	i := 0
	di := 0
	for _, char := range input {
		if di == 3 {
			di = 0
		}

		if di > 0 && char == ',' {
			i++
			di = 0
		}

		if pattern[i] == 'D' {
			if char >= 48 && char <= 57 {
				temp[i+di] = char
				di++
				continue
			}
		}

		if char == pattern[i] {
			temp[i] = char
			i++
		} else {
			clear(temp)
			i = 0
		}

		fmt.Println(string(temp))
	}

	fmt.Println(string(pattern))

	// temp := make([]byte, 12)
	// temp[0] = byte('a')
	// temp[1] = byte('b')
	// fmt.Println(string(temp))
}
