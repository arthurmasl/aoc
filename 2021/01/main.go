package main

import (
	"fmt"
	"strconv"

	"aoc/internal/utils"
)

func main() {
	lines := utils.GetLines("example")
	prevSum := 0
	increased := -1

	for chunk := range utils.Window(lines, 3) {
		sum := 0
		for _, n := range chunk {
			i, _ := strconv.Atoi(n)
			sum += i
		}

		if sum > prevSum {
			increased++
		}

		prevSum = sum
	}

	fmt.Println(increased)
}
