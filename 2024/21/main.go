package main

import (
	"fmt"
	"strconv"
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

func main() {
	inputs := utils.GetLines("input")

	total := 0
	for _, input := range inputs {
		seq1 := getSequence(numeric, numericPos, input)
		seq2 := getSequence(directional, directionalPos, seq1)
		seq3 := getSequence(directional, directionalPos, seq2)

		code, _ := strconv.Atoi(input[:len(input)-1])
		total += code * utf8.RuneCountInString(seq3)

		// fmt.Println(input)
		// fmt.Println(seq1)
		// fmt.Println(seq2)
		// fmt.Println(seq3)

		fmt.Println(utf8.RuneCountInString(seq3), code)
	}

	fmt.Println(total)
	// utils.Assert(total == 126384)
}

type keys struct {
	from rune
	to   rune
}

func getSequence(grid []string, initialPos vec, input string) string {
	seq := ""
	pos := initialPos

	for _, targetKey := range input {
		key := getKey(grid, pos)

		moves := paths[keys{key, targetKey}]
		seq += moves + "A"

		targetPos := getPos(grid, targetKey)
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

var paths = map[keys]string{
	{'A', '0'}: "<",
	{'0', 'A'}: ">",
	{'A', '1'}: "^<<",
	{'1', 'A'}: ">>v",
	{'A', '2'}: "<^",
	{'2', 'A'}: "v>",
	{'A', '3'}: "^",
	{'3', 'A'}: "v",
	{'A', '4'}: "^^<<",
	{'4', 'A'}: ">>vv",
	{'A', '5'}: "<^^",
	{'5', 'A'}: "vv>",
	{'A', '6'}: "^^",
	{'6', 'A'}: "vv",
	{'A', '7'}: "^^^<<",
	{'7', 'A'}: ">>vvv",
	{'A', '8'}: "<^^^",
	{'8', 'A'}: "vvv>",
	{'A', '9'}: "^^^",
	{'9', 'A'}: "vvv",
	{'0', '1'}: "^<",
	{'1', '0'}: ">v",
	{'0', '2'}: "^",
	{'2', '0'}: "v",
	{'0', '3'}: "^>",
	{'3', '0'}: "<v",
	{'0', '4'}: "^<^",
	{'4', '0'}: ">vv",
	{'0', '5'}: "^^",
	{'5', '0'}: "vv",
	{'0', '6'}: "^^>",
	{'6', '0'}: "<vv",
	{'0', '7'}: "^^^<",
	{'7', '0'}: ">vvv",
	{'0', '8'}: "^^^",
	{'8', '0'}: "vvv",
	{'0', '9'}: "^^^>",
	{'9', '0'}: "<vvv",
	{'1', '2'}: ">",
	{'2', '1'}: "<",
	{'1', '3'}: ">>",
	{'3', '1'}: "<<",
	{'1', '4'}: "^",
	{'4', '1'}: "v",
	{'1', '5'}: "^>",
	{'5', '1'}: "<v",
	{'1', '6'}: "^>>",
	{'6', '1'}: "<<v",
	{'1', '7'}: "^^",
	{'7', '1'}: "vv",
	{'1', '8'}: "^^>",
	{'8', '1'}: "<vv",
	{'1', '9'}: "^^>>",
	{'9', '1'}: "<<vv",
	{'2', '3'}: ">",
	{'3', '2'}: "<",
	{'2', '4'}: "<^",
	{'4', '2'}: "v>",
	{'2', '5'}: "^",
	{'5', '2'}: "v",
	{'2', '6'}: "^>",
	{'6', '2'}: "<v",
	{'2', '7'}: "<^^",
	{'7', '2'}: "vv>",
	{'2', '8'}: "^^",
	{'8', '2'}: "vv",
	{'2', '9'}: "^^>",
	{'9', '2'}: "<vv",
	{'3', '4'}: "<<^",
	{'4', '3'}: "v>>",
	{'3', '5'}: "<^",
	{'5', '3'}: "v>",
	{'3', '6'}: "^",
	{'6', '3'}: "v",
	{'3', '7'}: "<<^^",
	{'7', '3'}: "vv>>",
	{'3', '8'}: "<^^",
	{'8', '3'}: "vv>",
	{'3', '9'}: "^^",
	{'9', '3'}: "vv",
	{'4', '5'}: ">",
	{'5', '4'}: "<",
	{'4', '6'}: ">>",
	{'6', '4'}: "<<",
	{'4', '7'}: "^",
	{'7', '4'}: "v",
	{'4', '8'}: "^>",
	{'8', '4'}: "<v",
	{'4', '9'}: "^>>",
	{'9', '4'}: "<<v",
	{'5', '6'}: ">",
	{'6', '5'}: "<",
	{'5', '7'}: "<^",
	{'7', '5'}: "v>",
	{'5', '8'}: "^",
	{'8', '5'}: "v",
	{'5', '9'}: "^>",
	{'9', '5'}: "<v",
	{'6', '7'}: "<<^",
	{'7', '6'}: "v>>",
	{'6', '8'}: "<^",
	{'8', '6'}: "v>",
	{'6', '9'}: "^",
	{'9', '6'}: "v",
	{'7', '8'}: ">",
	{'8', '7'}: "<",
	{'7', '9'}: ">>",
	{'9', '7'}: "<<",
	{'8', '9'}: ">",
	{'9', '8'}: "<",
	{'<', '^'}: ">^",
	{'^', '<'}: "v<",
	{'<', 'v'}: ">",
	{'v', '<'}: "<",
	{'<', '>'}: ">>",
	{'>', '<'}: "<<",
	{'<', 'A'}: ">>^",
	{'A', '<'}: "v<<",
	{'^', 'v'}: "v",
	{'v', '^'}: "^",
	{'^', '>'}: "v>",
	{'>', '^'}: "<^",
	{'^', 'A'}: ">",
	{'A', '^'}: "<",
	{'v', '>'}: ">",
	{'>', 'v'}: "<",
	{'v', 'A'}: "^>",
	{'A', 'v'}: "<v",
	{'>', 'A'}: "^",
	{'A', '>'}: "v",
}
