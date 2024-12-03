package main

import (
	"fmt"
	"slices"
	"strings"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2023/09/input")

	total := 0
	for _, v := range lines {
		numString := strings.Split(v, " ")
		nums := utils.ConvertToInts(numString)

		levels := make([][]int, 0)
		levels = append(levels, nums)

		i := 0
		for {
			level := make([]int, 0)
			for chunk := range utils.Window(levels[i], 2) {
				diff := chunk[1] - chunk[0]
				level = append(level, diff)
			}
			levels = append(levels, level)

			if allZeros(level) {
				break
			}

			i++
		}

		// part 1
		// for _, level := range levels {
		// 	if len(level) == 0 {
		// 		break
		// 	}
		//
		// 	// fmt.Println(level, level[len(level)-1])
		// 	fmt.Println(level[0], level, level[len(level)-1])
		// 	total += level[len(level)-1]
		// }

		// part 2
		slices.Reverse(levels)
		left := 0
		for j, level := range levels {
			if len(levels) <= j+1 {
				break
			}
			left = levels[j+1][0] - left
			fmt.Println(left, level, j)
		}
		total += left
	}

	fmt.Println(total)
}

func allZeros(slice []int) bool {
	for _, value := range slice {
		if value != 0 {
			return false
		}
	}
	return true
}
