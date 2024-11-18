package main

import (
	"fmt"
	"strings"

	"aoc/src/internal/utils"
)

// func main() {
// 	lines := utils.GetLines("src/2020/06/input", "\n\n")
// 	sizes := []int{}
//
// 	for _, line := range lines {
// 		m := make(map[rune]bool)
// 		l := strings.ReplaceAll(strings.ReplaceAll(line, "\n", ""), " ", "")
//
// 		for _, r := range l {
// 			m[r] = true
// 		}
//
// 		sizes = append(sizes, len(m))
// 	}
//
// 	sum := 0
// 	for _, s := range sizes {
// 		sum += s
// 	}
//
// 	fmt.Println(sum)
// }

func main() {
	lines := utils.GetLines("src/2020/06/input", "\n\n")

	sizes := []int{}

	for _, line := range lines {
		groups := strings.Split(line, "\n")

		fmt.Println(runeIntersection(groups...))
		sizes = append(sizes, len(runeIntersection(groups...)))
	}

	sum := 0
	for _, s := range sizes {
		sum += s
	}

	fmt.Println(sum)
}

func runeIntersection(strings ...string) []string {
	intersection := make(map[rune]bool)
	for _, r := range strings[0] {
		intersection[r] = true
	}

	for _, str := range strings[1:] {
		current := make(map[rune]bool)

		for _, r := range str {
			if intersection[r] {
				current[r] = true
			}
		}

		intersection = current
	}

	var result []string
	for r := range intersection {
		result = append(result, string(r))
	}

	return result
}
