package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

// func main() {
// 	lines := utils.GetLines("src/2020/02/input")
//
// 	valid := 0
//
// 	for _, line := range lines {
// 		chunks := strings.Split(line, " ")
// 		numbers := strings.Split(chunks[0], "-")
// 		n1, _ := strconv.Atoi(numbers[0])
// 		n2, _ := strconv.Atoi(numbers[1])
// 		letter := string(chunks[1][0])
// 		word := chunks[2]
//
// 		letterCount := strings.Count(word, letter)
//
// 		fmt.Println(n1, n2, letterCount)
//
// 		if letterCount >= n1 && letterCount <= n2 {
// 			valid++
// 		}
// 	}
//
// 	fmt.Println(valid)
// }

func main() {
	lines := utils.GetLines("src/2020/02/input")

	valid := 0

	for _, line := range lines {
		chunks := strings.Split(line, " ")
		numbers := strings.Split(chunks[0], "-")
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		letter := string(chunks[1][0])
		word := chunks[2]

		fo := string(word[n1-1]) == letter
		so := string(word[n2-1]) == letter

		if fo != so {
			valid++
		}
	}

	fmt.Println(valid)
}
