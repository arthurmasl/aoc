package main

import (
	"container/list"
	"fmt"

	"aoc/src/internal/utils"
)

type Pos struct {
	x, y int
}

var directions = []Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	lines := utils.GetLines("example")

	startPos := Pos{}
	endPos := Pos{}

	visited := make(map[Pos]bool)
	parent := make(map[Pos]Pos)
	path := make([]Pos, 0)

	for y, row := range lines {
		for x, col := range row {
			if col == 'S' {
				startPos = Pos{x, y}
			}
			if col == 'E' {
				endPos = Pos{x, y}
			}
		}
	}

	queue := list.New()
	queue.PushBack(startPos)

	visited[startPos] = true

	for queue.Len() > 0 {
		element := queue.Front()
		currentPos := element.Value.(Pos)

		queue.Remove(element)

		if currentPos == endPos {
			fmt.Println("end found")
			for currentPos != startPos {
				path = append(path, currentPos)
				currentPos = parent[currentPos]
			}
			break
		}

		for _, neighbor := range getNeighbors(lines, currentPos) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				parent[neighbor] = currentPos

			}
		}
	}

	// fill path
	for _, pos := range path {
		for y, row := range lines {
			newRow := []byte(row)
			for x := range row {
				if pos.x == x && pos.y == y {
					newRow[x] = '@'
				}
			}
			lines[y] = string(newRow)
		}
	}

	// draw grid
	for _, row := range lines {
		fmt.Println(row)
	}
}

func getNeighbors(lines []string, pos Pos) []Pos {
	neighbors := make([]Pos, 0)

	for _, dir := range directions {
		neighborPos := Pos{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(lines, neighborPos.x, neighborPos.y)

		if ok && neighborId != '#' {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}
