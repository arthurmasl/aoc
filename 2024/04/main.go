package main

import (
	"fmt"
	"strings"

	"aoc/internal/utils"
)

//  M M    S S    S M     M S       S M
//   A      A      A       A         A    <-bad
//  S S    M M    S M     M S       M S

func main() {
	lines := utils.GetLines("src/2024/04/input")

	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}

	total := 0

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 'A' {
				tl := safeGet(matrix, y-1, x-1)
				tr := safeGet(matrix, y-1, x+1)
				bl := safeGet(matrix, y+1, x-1)
				br := safeGet(matrix, y+1, x+1)

				if tl == br || tr == bl {
					continue
				}

				hm := make(map[string]int)
				hm[tl] += 1
				hm[tr] += 1
				hm[bl] += 1
				hm[br] += 1

				if hm["S"] == 2 && hm["M"] == 2 {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func safeGet(arr [][]string, row, col int) string {
	if row >= 0 && row < len(arr) && col >= 0 && col < len(arr[row]) {
		return arr[row][col]
	}
	return ""
}
