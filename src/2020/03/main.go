package main

import (
	"fmt"

	"aoc/src/internal/utils"
)

var (
	x = 0
	y = 0
	l = 0
)

func move(mx, my int) {
	x += mx
	y += my

	if x >= l {
		diff := l - x
		x = diff - diff*2
	}
}

func main() {
	lines := utils.GetLines("src/2020/03/input")
	l = len(lines[0])
	routes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	treeCounts := []int{}

	for _, route := range routes {
		trees := 0
		x = 0
		y = 0

		for {
			move(route[0], route[1])

			if y >= len(lines) {
				treeCounts = append(treeCounts, trees)
				break
			}

			if string(lines[y][x]) == "#" {
				trees++
			}

		}
	}

	sum := 0
	for _, s := range treeCounts {
		if sum == 0 {
			sum = s
			continue
		}
		sum *= s
	}

	fmt.Println(sum)
}
