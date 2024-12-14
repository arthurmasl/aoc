package main

import (
	"fmt"
	"strings"

	"aoc/src/internal/utils"
)

func main() {
	line := utils.GetLines("example")[0]
	numbers := utils.ConvertToInts(strings.Split(line, ","))

	fishes := make(map[int]int)
	for _, n := range numbers {
		fishes[n]++
	}

	for range 256 {
		newFishes := make(map[int]int)
		for k, v := range fishes {
			newFishes[k] = v
		}

		for timer, count := range fishes {
			if count == 0 {
				continue
			}
			if timer == 0 {
				newFishes[0] -= count
				newFishes[6] += count
				newFishes[8] += count
				continue
			}

			newFishes[timer] -= count
			newFishes[timer-1] += count
		}

		fishes = newFishes
	}

	// fmt.Println(fishes)
	sum := 0
	for _, count := range fishes {
		sum += count
		// for range count {
		// 	fmt.Print(timer)
		// 	fmt.Print(" ")
		// }
	}
	fmt.Println(sum)
}
