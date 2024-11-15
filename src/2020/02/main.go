package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2020/02/input")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func parse(line string) (int, int, string, string) {
	chunks := strings.Split(line, " ")
	numbers := strings.Split(chunks[0], "-")
	n1, _ := strconv.Atoi(numbers[0])
	n2, _ := strconv.Atoi(numbers[1])
	letter := string(chunks[1][0])
	word := chunks[2]

	return n1, n2, letter, word
}

func part1(lines []string) int {
	valid := 0

	for _, line := range lines {
		n1, n2, letter, word := parse(line)
		letterCount := strings.Count(word, letter)

		if letterCount >= n1 && letterCount <= n2 {
			valid++
		}
	}

	return valid
}

func part2(lines []string) int {
	valid := 0

	for _, line := range lines {
		n1, n2, letter, word := parse(line)

		fo := string(word[n1-1]) == letter
		so := string(word[n2-1]) == letter

		if fo != so {
			valid++
		}
	}

	return valid
}
