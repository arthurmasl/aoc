package main

import (
	"fmt"
	"math"

	"aoc/src/internal/utils"
)

type Vector struct {
	x, y int
}

type Node struct {
	pos Vector
	dir rune
}

type State struct {
	node Node
	path []Vector
	cost int
}

var (
	directions = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	arrows     = []rune{'^', '>', 'v', '<'}
)

func main() {
	grid := utils.GetLines("example")
	start, end := getInitialPoints(grid)

	lowestCost := math.MaxInt
	visited := make(map[Node]int)
	paths := make(map[int][]Vector)

	queue := []State{{node: Node{start, arrows[1]}, path: []Vector{start}, cost: 0}}

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		if currentState.cost > lowestCost {
			continue
		}

		if currentState.node.pos == end {
			if currentState.cost <= lowestCost {
				lowestCost = currentState.cost
				paths[lowestCost] = append(paths[lowestCost], currentState.path...)
			}
			continue
		}

		for _, neighbor := range getNeighbors(grid, currentState.node.pos) {
			newCost := currentState.cost + 1
			if currentState.node.dir != neighbor.dir {
				newCost += 1000
			}

			if cost, ok := visited[neighbor]; !ok || newCost <= cost {
				visited[neighbor] = newCost

				queue = append(queue, State{
					node: neighbor,
					path: append(append([]Vector{}, currentState.path...), neighbor.pos),
					cost: newCost,
				})
			}
		}
	}

	for _, pos := range paths[lowestCost] {
		for y, row := range grid {
			newRow := []byte(row)
			for x := range row {
				if pos.x == x && pos.y == y {
					newRow[x] = 'O'
				}
			}
			grid[y] = string(newRow)
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	total := 0
	for _, row := range grid {
		for _, col := range row {
			if col == 'O' {
				total++
			}
		}
	}

	fmt.Println(total)
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

func getNeighbors(lines []string, pos Vector) []Node {
	neighbors := make([]Node, 0)

	for i, dir := range directions {
		neighborPos := Vector{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(lines, neighborPos.x, neighborPos.y)

		if ok && neighborId != '#' {
			neighbors = append(neighbors, Node{neighborPos, arrows[i]})
		}
	}

	return neighbors
}
