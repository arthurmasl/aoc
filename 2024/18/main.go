package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

// const (
// 	size  = 70
// 	bytes = 1024
// )

const (
	size  = 6
	bytes = 12
)

type Vector struct {
	x, y int
}

var directions = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	walls := utils.GetLines("example")[:bytes]
	grid := make([]string, size+1)

	start := Vector{0, 0}
	end := Vector{size, size}

	for y := range grid {
		newRow := strings.Repeat(".", size+1)
		grid[y] = newRow
	}

	for _, wall := range walls {
		xStr, yStr, _ := strings.Cut(wall, ",")
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)

		newRow := []byte(grid[y])
		newRow[x] = '#'
		grid[y] = string(newRow)
	}

	path := make([]Vector, 0)

	visited := make(map[Vector]bool)
	visited[start] = true
	parent := make(map[Vector]Vector)

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Front()
		current := node.Value.(Vector)
		queue.Remove(node)

		if current == end {
			path = make([]Vector, 0)
			for current != start {
				path = append(path, current)
				current = parent[current]
			}

			fmt.Println("found", len(path))
			continue
		}

		for _, neighbor := range getNeighbors(grid, current) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}

	for _, pos := range path {
		newRow := []byte(grid[pos.y])
		newRow[pos.x] = 'O'
		grid[pos.y] = string(newRow)
		// }
	}

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }
}

func getNeighbors(grid []string, pos Vector) []Vector {
	neighbors := make([]Vector, 0)

	for _, dir := range directions {
		neighborPos := Vector{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(grid, neighborPos.x, neighborPos.y)

		if ok && neighborId != '#' {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}
