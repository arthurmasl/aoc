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
// 		letter := strings.Split(strings.Split(line, ": ")[0], " ")[1]
// 		word := strings.Split(line, ": ")[1]
// 		numbers := strings.Split(strings.Split(line, " ")[0], "-")
// 		from, _ := strconv.Atoi(numbers[0])
// 		to, _ := strconv.Atoi(numbers[1])
//
// 		letterCount := strings.Count(word, letter)
//
// 		fmt.Println(from, to, letterCount)
//
// 		if letterCount >= from && letterCount <= to {
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
		letter := strings.Split(strings.Split(line, ": ")[0], " ")[1]
		word := strings.Split(line, ": ")[1]
		numbers := strings.Split(strings.Split(line, " ")[0], "-")
		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])

		fmt.Println(word, first, second, word[first], word[second-1])

		fo := string(word[first-1]) == letter
		so := string(word[second-1]) == letter

		if fo != so {
			valid++
		}
	}

	fmt.Println(valid)
}
