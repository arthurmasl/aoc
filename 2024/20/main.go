package main

import (
	"container/list"
	"fmt"

	"aoc/internal/utils"
)

type Vector struct {
	x, y int
}

var (
	directions = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	arrows     = []rune{'^', '>', 'v', '<'}
)

func main() {
	grid := utils.GetLines("input")
	start, end := getInitialPoints(grid)
	walls := getWalls(grid)

	defaultPathLen := getPathLen(grid, start, end)

	checked := 0
	count := 0
	for _, wall := range walls {
		newGrid := make([]string, len(grid))
		copy(newGrid, grid)
		newRow := []byte(grid[wall.y])
		newRow[wall.x] = '.'
		newGrid[wall.y] = string(newRow)

		newPathLen := getPathLen(newGrid, start, end)
		diff := defaultPathLen - newPathLen

		checked++
		fmt.Println(checked, "of", len(walls))
		if diff >= 100 {
			count++
		}
	}

	fmt.Println(count)

	// for _, pos := range paths[lowestCost] {
	// 	for y, row := range grid {
	// 		newRow := []byte(row)
	// 		for x := range row {
	// 			if pos.x == x && pos.y == y {
	// 				newRow[x] = 'o'
	// 			}
	// 		}
	// 		grid[y] = string(newRow)
	// 	}
	// }

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }
}

func getPathLen(grid []string, start, end Vector) int {
	distance := make(map[Vector]int)
	distance[start] = 0

	visited := make(map[Vector]bool)
	visited[start] = true

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Front()
		current := node.Value.(Vector)
		queue.Remove(node)

		if current == end {
			return distance[current]
		}

		for _, neighbor := range getNeighbors(grid, current) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				distance[neighbor] = distance[current] + 1
			}
		}
	}

	return -1
}

func getInitialPoints(grid []string) (Vector, Vector) {
	var start, end Vector

	for y, row := range grid {
		for x, col := range row {
			if col == 'S' {
				start = Vector{x, y}
			}
			if col == 'E' {
				end = Vector{x, y}
			}
		}
	}

	return start, end
}

func getNeighbors(grid []string, pos Vector) []Vector {
	neighbors := make([]Vector, 0)

	for _, dir := range directions {
		neighborPos := Vector{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(grid, neighborPos.x, neighborPos.y)

		if ok && neighborId != '#' && neighborId != 'W' {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}

func getWalls(grid []string) []Vector {
	walls := make([]Vector, 0)

	for y, row := range grid {
		for x, col := range row {
			if col == '#' {
				walls = append(walls, Vector{x, y})
			}
		}
	}

	return walls
}
