package main

import (
	"fmt"
	"math"

	"aoc/internal/utils"
)

type vec struct {
	x, y int
}

var (
	numericPos = vec{2, 3}
	numeric    = []string{
		"789",
		"456",
		"123",
		"_0A",
	}
)

var (
	directionalPos = vec{2, 0}
	directional    = []string{
		"_^A",
		"<v>",
	}
)

func main() {
	inputs := utils.GetLines("example")

	utils.Assert(string(numeric[numericPos.y][numericPos.x]) == "A")
	utils.Assert(string(directional[directionalPos.y][directionalPos.x]) == "A")

	seq1 := getSequence(numeric, numericPos, inputs[0])
	fmt.Println(seq1)
	utils.Assert(getSequence(numeric, numericPos, inputs[0]) == "<A^A>^^AvvvA")
}

func getSequence(grid []string, initialPos vec, input string) string {
	seq := ""
	pos := initialPos

	for _, targetKey := range input {
		targetPos := getPos(grid, targetKey)
		deltaX := targetPos.x - pos.x
		deltaY := targetPos.y - pos.y

		if deltaX < 0 {
			for range int(math.Abs(float64(deltaX))) {
				seq += "<"
			}
		} else {
			for range deltaX {
				seq += ">"
			}
		}

		if deltaY < 0 {
			for range int(math.Abs(float64(deltaY))) {
				seq += "^"
			}
		} else {
			for range deltaY {
				seq += "v"
			}
		}

		seq += "A"
		pos = targetPos
	}

	return seq
}

func getKey(grid []string, pos vec) rune {
	return rune(grid[pos.y][pos.x])
}

func getPos(grid []string, key rune) vec {
	for y, row := range grid {
		for x, col := range row {
			if col == key {
				return vec{x, y}
			}
		}
	}

	return vec{}
}
