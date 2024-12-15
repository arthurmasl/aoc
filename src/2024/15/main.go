package main

import (
	"fmt"
	"strings"

	"aoc/src/internal/utils"
)

type Pos struct {
	x, y int
}

func main() {
	blocks := utils.GetLines("example", "\n\n")
	grid := strings.Split(blocks[0], "\n")
	robotPos := Pos{}
	moves := blocks[1]

	// update grid
	for y, row := range grid {
		newRow := ""

		for _, col := range row {
			switch col {
			case '#':
				newRow += "##"
			case 'O':
				newRow += "[]"
			case '.':
				newRow += ".."
			case '@':
				newRow += "@."
			}
		}

		grid[y] = newRow
	}

	// find robot
	for y, row := range grid {
		for x, col := range row {
			if col == '@' {
				robotPos = Pos{x, y}
			}
		}
	}

	// move
	for _, move := range moves {
		tryToMove(&grid, move, robotPos, &robotPos)
	}

	// draw grid
	for _, row := range grid {
		fmt.Println(row)
	}

	// calculate result
	sum := 0
	for y, row := range grid {
		for x, col := range row {
			if col == 'O' {
				sum += 100*y + x
			}
		}
	}

	fmt.Println(sum)
}

func tryToMove(grid *[]string, move rune, pos Pos, robotPos *Pos) bool {
	vel := getVel(move)
	nextPos := applyVel(pos, vel)
	nextCell, _ := utils.GetSafeValue(*grid, nextPos.x, nextPos.y)

	switch nextCell {
	case '.':
		moveElement(grid, pos, nextPos)
		*robotPos = nextPos
		return true
	case 'O':
		ok := tryToMove(grid, move, nextPos, robotPos)
		if ok {
			tryToMove(grid, move, pos, robotPos)
			return ok
		}
	}

	return false
}

func moveElement(grid *[]string, from, to Pos) {
	symbolToMove := rune((*grid)[from.y][from.x])

	toRow := []rune((*grid)[to.y])
	toRow[to.x] = symbolToMove
	(*grid)[to.y] = string(toRow)

	fromRow := []rune((*grid)[from.y])
	fromRow[from.x] = '.'
	(*grid)[from.y] = string(fromRow)
}

func getVel(symbol rune) Pos {
	pos := Pos{}

	switch symbol {
	case '<':
		pos.x = -1
	case '>':
		pos.x = 1
	case '^':
		pos.y = -1
	case 'v':
		pos.y = 1
	}

	return pos
}

func applyVel(pos Pos, vel Pos) Pos {
	pos.x += vel.x
	pos.y += vel.y
	return pos
}
