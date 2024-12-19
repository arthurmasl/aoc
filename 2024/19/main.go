package main

import (
	"fmt"
	"slices"
	"strings"

	"aoc/internal/utils"
)

func main() {
	blocks := utils.GetLines("input", "\n\n")
	patterns := strings.Split(blocks[0], ", ")
	designs := strings.Split(blocks[1], "\n")

	possibleDesigns := make(map[string]bool)

	for _, design := range designs {
		target := strings.Repeat("x", len(design))
		draft := design

		for _, pattern := range slices.SortedFunc(slices.Values(patterns), sortByLength) {
			draft = strings.ReplaceAll(draft, pattern, strings.Repeat("x", len(pattern)))
			if draft == target {
				possibleDesigns[design] = true
				break
			}
		}

	}

	fmt.Println(len(possibleDesigns))
}

func sortByLength(a, b string) int {
	return len(b) - len(a)
}

func sortByLengthRev(a, b string) int {
	return len(a) - len(b)
}
