package main

import (
	"container/list"
	"fmt"
	"slices"

	"aoc/internal/utils"
)

type Vector struct {
	x, y int
}

var (
	directions = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	arrows     = []rune{'^', '>', 'v', '<'}
)

const cheatTick = 20

func main() {
	grid := utils.GetLines("input")
	emptyGrid := getEmptyGrid(grid)
	start, end := getInitialPoints(grid)
	// walls := getWalls(grid)

	defaultPath, distances := getPath(grid, start, end)

	count := 0
	for i, cheatStartPath := range defaultPath {
		availablePaths := getAvailablePaths(grid, cheatStartPath, cheatTick)

		for _, cheatEndPath := range availablePaths {
			_, cheatDistances := getPath(emptyGrid, cheatStartPath, cheatEndPath)
			cheatSteps := cheatDistances[cheatEndPath]

			newDistance := distances[end] - distances[cheatEndPath] + distances[cheatStartPath] + cheatSteps
			diff := distances[end] - newDistance

			if diff >= 100 {
				count++
			}

		}

		fmt.Println(i, "of", len(defaultPath))

	}
	fmt.Println(count)

	// for _, pos := range availablePaths {
	// 	for y, row := range grid {
	// 		newRow := []byte(row)
	// 		for x, col := range row {
	// 			if col != 'S' && pos.x == x && pos.y == y {
	// 				newRow[x] = 'x'
	// 			}
	// 		}
	// 		grid[y] = string(newRow)
	// 	}
	// }
	//
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }
}

func getPath(grid []string, start, end Vector) ([]Vector, map[Vector]int) {
	path := make([]Vector, 0)
	parent := make(map[Vector]Vector)

	distances := make(map[Vector]int)
	distances[start] = 0

	visited := make(map[Vector]bool)
	visited[start] = true

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
			path = append(path, start)
			slices.Reverse(path)
			return path, distances
		}

		for _, neighbor := range getNeighbors(grid, current) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
				distances[neighbor] = distances[current] + 1
			}
		}
	}

	return path, distances
}

func getAvailablePaths(grid []string, pos Vector, size int) []Vector {
	paths := make([]Vector, 0)

	for y := -size; y <= 0; y++ {
		for x := -size - y; x <= size+y; x++ {
			path := Vector{x + pos.x, y + pos.y}
			pathId, ok := utils.GetSafeValue(grid, path.x, path.y)

			if ok && pathId != '#' {
				paths = append(paths, path)
			}
		}
	}

	for y := 1; y <= size; y++ {
		for x := -size + y; x <= size-y; x++ {
			path := Vector{x + pos.x, y + pos.y}
			pathId, ok := utils.GetSafeValue(grid, path.x, path.y)

			if ok && pathId != '#' {
				paths = append(paths, path)
			}
		}
	}

	return paths
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

func getEmptyGrid(grid []string) []string {
	newGrid := make([]string, len(grid))
	copy(newGrid, grid)

	for y, row := range newGrid {
		newRow := []byte(row)
		for x, col := range row {
			if col == '#' {
				newRow[x] = '.'
			}
		}

		newGrid[y] = string(newRow)
	}

	return newGrid
}
