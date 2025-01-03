package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

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

// 133676 to high
// 132532
// 132312 to low
func main() {
	inputs := utils.GetLines("input")

	total := 0
	for _, input := range inputs {
		seq1 := getSequence(numeric, numericPos, input)
		seq2 := getSequence(directional, directionalPos, seq1)
		seq3 := getSequence(directional, directionalPos, seq2)

		code, _ := strconv.Atoi(input[:len(input)-1])
		total += code * utf8.RuneCountInString(seq3)

		fmt.Println(input)
		fmt.Println(seq1)
		fmt.Println(seq2)
		fmt.Println(seq3)

		fmt.Println(utf8.RuneCountInString(seq3), code)
	}

	fmt.Println(total)
	// utils.Assert(total == 126384)
}

func getSequence(grid []string, initialPos vec, input string) string {
	seq := ""
	pos := initialPos

	for _, targetKey := range input {
		targetPos := getPos(grid, targetKey)
		moves := getMoves(grid, pos, targetPos)

		seq += moves + "A"
		fmt.Println(seq)
		pos = targetPos
	}

	return seq
}

func getMoves(grid []string, pos, targetPos vec) string {
	moves := make([]rune, 0)
	deltaX := targetPos.x - pos.x
	deltaY := targetPos.y - pos.y

	keys := ""

	moveX(grid, deltaX, &moves, &pos, &keys)
	moveY(grid, deltaY, &moves, &pos, &keys)

	hasBlank := strings.Contains(keys, "_")
	if hasBlank {
		slices.Reverse(moves)
	}

	return string(moves)
}

func moveX(grid []string, deltaX int, moves *[]rune, pos *vec, keys *string) {
	if deltaX < 0 {
		for range deltaX * -1 {
			*moves = append(*moves, '<')
			pos.x -= 1
			*keys += string(getKey(grid, vec{pos.x, pos.y}))
		}
	} else {
		for range deltaX {
			*moves = append(*moves, '>')
			pos.x += 1
			*keys += string(getKey(grid, vec{pos.x, pos.y}))
		}
	}
}

func moveY(grid []string, deltaY int, moves *[]rune, pos *vec, keys *string) {
	if deltaY < 0 {
		for range deltaY * -1 {
			*moves = append(*moves, '^')
			pos.y -= 1
			*keys += string(getKey(grid, vec{pos.x, pos.y}))
		}
	} else {
		for range deltaY {
			*moves = append(*moves, 'v')
			pos.y += 1
			*keys += string(getKey(grid, vec{pos.x, pos.y}))
		}
	}
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
