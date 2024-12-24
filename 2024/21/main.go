package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"

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

// 132312 to low
// 20592 to low
func main() {
	inputs := utils.GetLines("input")

	total := 0
	for _, input := range inputs {
		seq1 := getSequence(numeric, numericPos, input)
		seq2 := getSequence(directional, directionalPos, seq1)
		seq3 := getSequence(directional, directionalPos, seq2)

		code, _ := strconv.Atoi(input[:len(input)-1])
		total += code * len(seq3)

		// fmt.Println(input)
		// fmt.Println(seq1)
		// fmt.Println(seq2)
		// fmt.Println(seq3)

		fmt.Println(len(seq3), code)
	}

	fmt.Println(total)
	// utils.Assert(total == 126384)
}

type mov struct {
	pos vec
	key rune
}

func getSequence(grid []string, initialPos vec, input string) string {
	seq := ""
	pos := initialPos

	for _, targetKey := range input {
		targetPos := getPos(grid, targetKey)
		moves := getMoves(pos, targetPos)

		hasBlank := false
		for _, move := range moves {
			if getKey(grid, move.pos) == '_' {
				hasBlank = true
			}
		}

		if hasBlank {
			slices.Reverse(moves)
		}

		for _, move := range moves {
			seq += string(move.key)
			pos = move.pos
		}

		seq += "A"
		pos = targetPos
	}

	return seq
}

func getMoves(pos, targetPos vec) []mov {
	deltaX := targetPos.x - pos.x
	deltaY := targetPos.y - pos.y
	moves := make([]mov, 0)

	if deltaX < 0 {
		for x := range int(math.Abs(float64(deltaX))) {
			moves = append(moves, mov{vec{pos.x - (x + 1), pos.y}, '<'})
		}
	} else {
		for x := range deltaX {
			moves = append(moves, mov{vec{pos.x + (x + 1), pos.y}, '>'})
		}
	}

	if deltaY < 0 {
		for y := range int(math.Abs(float64(deltaY))) {
			moves = append(moves, mov{vec{pos.x, pos.y - (y + 1)}, '^'})
		}
	} else {
		for y := range deltaY {
			moves = append(moves, mov{vec{pos.x, pos.y + (y + 1)}, 'v'})
		}
	}

	return moves
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
