package main

import (
	"fmt"
	"strings"

	"aoc/internal/utils"
)

func main() {
	blocks := utils.GetLines("input", "\n\n")
	patterns := strings.Split(blocks[0], ", ")
	designs := strings.Split(blocks[1], "\n")

	memo := make(map[string]int)

	var getPossibleWays func(string) int
	getPossibleWays = func(design string) int {
		ways, ok := memo[design]
		if ok {
			return ways
		}

		if design == "" {
			return 1
		}

		for _, pattern := range patterns {
			if strings.HasPrefix(design, pattern) {
				ways += getPossibleWays(design[len(pattern):])
			}
		}

		memo[design] = ways
		return ways
	}

	ways := 0
	for _, design := range designs {
		ways += getPossibleWays(design)
	}

	fmt.Println(ways)
}
