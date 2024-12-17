package main

import (
	"fmt"
	"strings"

	"aoc/internal/utils"
)

type Pos struct {
	x, y int
}

func main() {
	blocks := utils.GetLines("example", "\n\n")
	grid := strings.Split(blocks[0], "\n")
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

	// move
	for _, move := range moves {
		prevGrid := make([]string, len(grid))
		copy(prevGrid, grid)

		movesList = nil
		robotPos := getRobotPos(grid)
		tryToMove(&grid, move, robotPos)

		revertMoves := false
		for _, move := range movesList {
			if !move.isGood {
				revertMoves = true
				break
			}
		}

		if revertMoves {
			grid = prevGrid
		}

	}

	// draw grid
	for _, row := range grid {
		fmt.Println(row)
	}

	// calculate result
	sum := 0
	for y, row := range grid {
		for x, col := range row {
			if col == '[' {
				sum += 100*y + x
			}
		}
	}

	fmt.Println(sum)
}

type Move struct {
	move   rune
	pos    Pos
	isGood bool
}

var movesList = make([]Move, 0)

func tryToMove(grid *[]string, move rune, pos Pos) bool {
	vel := getVel(move)
	nextPos := applyVel(pos, vel)
	nextCell, _ := utils.GetSafeValue(*grid, nextPos.x, nextPos.y)

	switch nextCell {
	case '.':
		moveElement(grid, pos, nextPos)
		return true

	case '[', ']':
		if move == '^' || move == 'v' {
			var nextSiblingPos Pos
			switch nextCell {
			case ']':
				nextSiblingPos = Pos{pos.x - 1, nextPos.y}
			case '[':
				nextSiblingPos = Pos{pos.x + 1, nextPos.y}
			}

			okLeft := tryToMove(grid, move, nextPos)

			if okLeft {
				okRight := tryToMove(grid, move, nextSiblingPos)
				if !okRight {
					// rollback left
					movedLeft := applyVel(nextPos, getVel(move))
					moveElement(grid, movedLeft, nextPos)
					movesList = append(movesList, Move{getReverseMove(move), movedLeft, false})
					return false
				}

				if okRight {
					ok := tryToMove(grid, move, pos)
					movesList = append(movesList, Move{move, pos, ok})
					return ok
				}
			}
		} else {
			if ok := tryToMove(grid, move, nextPos); ok {
				ok := tryToMove(grid, move, pos)
				movesList = append(movesList, Move{move, pos, ok})
				return ok
			}
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

func getReverseMove(symbol rune) rune {
	switch symbol {
	case '<':
		return '>'
	case '>':
		return '<'
	case '^':
		return 'v'
	case 'v':
		return '^'
	}
	return 0
}

func applyVel(pos Pos, vel Pos) Pos {
	pos.x += vel.x
	pos.y += vel.y
	return pos
}

func getRobotPos(grid []string) Pos {
	robotPos := Pos{}

	for y, row := range grid {
		for x, col := range row {
			if col == '@' {
				robotPos = Pos{x, y}
			}
		}
	}

	return robotPos
}
