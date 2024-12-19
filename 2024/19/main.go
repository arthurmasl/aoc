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

	memo := make(map[string]bool)

	var isPossible func(string) bool
	isPossible = func(design string) bool {
		has, ok := memo[design]
		if ok {
			return true
		}

		if design == "" {
			return true
		}

		for _, pattern := range patterns {
			if strings.HasPrefix(design, pattern) {
				has = isPossible(design[len(pattern):])
			}
		}

		memo[design] = has
		return has
	}

	count := 0
	for _, design := range designs {
		if isPossible(design) {
			count++
		}
	}

	fmt.Println(count)
}
