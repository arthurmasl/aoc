package main

import (
	"fmt"
	"slices"

	"aoc/src/internal/utils"
)

type Pos struct {
	x, y int
}

var visited = make(map[Pos]bool)

// 3916431 to high
// 874986 to high
// 830523 no
// 825587 no
// 747403 no
// 690001 to low
func main() {
	lines := utils.GetLines("input")
	total := 0

	for y, row := range lines {
		for x, col := range row {
			pos := Pos{x, y}
			if _, ok := visited[pos]; ok {
				continue
			}

			name := string(col)
			positions := make([]Pos, 0)
			searchNext(lines, &positions, pos)
			// fmt.Println(name, positions)

			area := len(positions)

			corners := 0

			emptyVisited := make(map[Pos]bool)
			for _, pos := range positions {
				// lt
				if getAt(lines, Pos{pos.x - 1, pos.y}) != name &&
					getAt(lines, Pos{pos.x, pos.y - 1}) != name {
					corners += 1
				}
				// rt
				if getAt(lines, Pos{pos.x + 1, pos.y}) != name &&
					getAt(lines, Pos{pos.x, pos.y - 1}) != name {
					corners += 1
				}
				// rb
				if getAt(lines, Pos{pos.x + 1, pos.y}) != name &&
					getAt(lines, Pos{pos.x, pos.y + 1}) != name {
					corners += 1
				}
				// lb
				if getAt(lines, Pos{pos.x - 1, pos.y}) != name &&
					getAt(lines, Pos{pos.x, pos.y + 1}) != name {
					corners += 1
				}

				if getAt(
					lines,
					Pos{pos.x + 1, pos.y + 1},
				) == getAt(
					lines,
					Pos{pos.x - 1, pos.y - 1},
				) {
					continue
				}
				// if getAt(
				// 	lines,
				// 	Pos{pos.x - 1, pos.y + 1},
				// ) == getAt(
				// 	lines,
				// 	Pos{pos.x + 1, pos.y - 1},
				// ) {
				// 	continue
				// }

				// if getAt(lines, Pos{pos.x + 1, pos.y + 1}) != name ||
				// 	getAt(lines, Pos{pos.x - 1, pos.y - 1}) != name {
				// 	continue
				// }

				emptyAround := getEmptyAround(lines, pos)
				for _, emptyPos := range emptyAround {
					if emptyVisited[emptyPos] {
						continue
					}

					if slices.Contains(positions, emptyPos) {
						continue
					}

					// if slices.Contains(positions, Pos{emptyPos.x + 1, emptyPos.y + 1}) {
					// 	continue
					// }
					// if slices.Contains(positions, Pos{emptyPos.x - 1, emptyPos.y - 1}) {
					// 	continue
					// }

					// lt
					if slices.Contains(positions, Pos{emptyPos.x - 1, emptyPos.y}) &&
						slices.Contains(positions, Pos{emptyPos.x, emptyPos.y - 1}) {
						corners += 1
						emptyVisited[emptyPos] = true
					}
					// rt
					if slices.Contains(positions, Pos{emptyPos.x + 1, emptyPos.y}) &&
						slices.Contains(positions, Pos{emptyPos.x, emptyPos.y - 1}) {
						corners += 1
						emptyVisited[emptyPos] = true
					}
					// rb
					if slices.Contains(positions, Pos{emptyPos.x + 1, emptyPos.y}) &&
						slices.Contains(positions, Pos{emptyPos.x, emptyPos.y + 1}) {
						corners += 1
						emptyVisited[emptyPos] = true
					}
					// lb
					if slices.Contains(positions, Pos{emptyPos.x - 1, emptyPos.y}) &&
						slices.Contains(positions, Pos{emptyPos.x, emptyPos.y + 1}) {
						corners += 1
						emptyVisited[emptyPos] = true
					}
				}

			}

			fmt.Println(name, corners)
			// perimeter = len(sides)
			price := area * corners
			total += price
		}
	}

	fmt.Println(total)
}

func getEmptyAround(lines []string, pos Pos) []Pos {
	positions := make([]Pos, 0)

	top := Pos{pos.x, pos.y - 1}
	right := Pos{pos.x + 1, pos.y}
	bottom := Pos{pos.x, pos.y + 1}
	left := Pos{pos.x - 1, pos.y}

	tryPositions := []Pos{top, right, bottom, left}
	for _, p := range tryPositions {
		if hasPos(lines, p) {
			positions = append(positions, p)
		}
	}

	return positions
}

func getAt(arr []string, pos Pos) string {
	if pos.y >= 0 && pos.y < len(arr) && pos.x >= 0 && pos.x < len(arr[pos.y]) {
		return string(arr[pos.y][pos.x])
	}
	return ""
}

func hasPos(arr []string, pos Pos) bool {
	if pos.y >= 0 && pos.y < len(arr) && pos.x >= 0 && pos.x < len(arr[pos.y]) {
		return true
	}
	return false
}

func searchNext(list []string, items *[]Pos, pos Pos) {
	if visited[pos] {
		return
	}

	curr := list[pos.y][pos.x]
	visited[pos] = true

	(*items) = append(*items, pos)

	checNext := func(y, x int) {
		if y >= 0 && y < len(list) && x >= 0 && x < len(list[y]) {
			if curr == list[y][x] {
				searchNext(list, items, Pos{x, y})
			}
		}
	}

	checNext(pos.y, pos.x-1)
	checNext(pos.y, pos.x+1)
	checNext(pos.y-1, pos.x)
	checNext(pos.y+1, pos.x)
}
