package main

import (
	"aoc/src/internal/utils"
	"fmt"
)

type Pos struct {
	x, y int
}

var visited = make(map[Pos]bool)

func main() {
	lines := utils.GetLines("src/2024/12/input")
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
			perimeter := 0

			sides := make(map[int]int)
			for _, position := range positions {
				t := Pos{position.x, position.y - 1}
				b := Pos{position.x, position.y + 1}
				l := Pos{position.x - 1, position.y}
				r := Pos{position.x + 1, position.y}

				// perimeter++
				if getAt(lines, t) != name {
					perimeter++
				}
				if getAt(lines, b) != name {
					perimeter++
				}
				if getAt(lines, l) != name {
					perimeter++
				}
				if getAt(lines, r) != name {
					perimeter++
				}
			}

			// perimeter = len(sides)
			price := area * perimeter

			if name == "A" {
				// fmt.Println(name, area, perimeter)
				fmt.Println(sides)
				fmt.Println(positions, perimeter)
			}
			total += price
		}
	}

	fmt.Println(total)
}

func getAt(arr []string, pos Pos) string {
	if pos.y >= 0 && pos.y < len(arr) && pos.x >= 0 && pos.x < len(arr[pos.y]) {
		return string(arr[pos.y][pos.x])
	}
	return ""
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
