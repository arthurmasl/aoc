package main

import (
	"container/list"
	"fmt"
	"strconv"

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

// 20592 to low
func main() {
	inputs := utils.GetLines("example")

	total := 0
	for _, input := range inputs[0:1] {
		seq1 := getSequence(numeric, numericPos, input)
		seq2 := getSequence(directional, directionalPos, seq1)
		seq3 := getSequence(directional, directionalPos, seq2)

		code, _ := strconv.Atoi(input[:len(input)-1])
		total += code * len(seq3)

		fmt.Println(input)
		fmt.Println(seq1)
		fmt.Println(seq2)
		fmt.Println(seq3)

		fmt.Println(len(seq3), code)
	}

	fmt.Println(total)
	utils.Assert(total == 126384)
}

type mov struct {
	pos vec
	key rune
}

func getSequence(grid []string, initialPos vec, input string) string {
	seq := ""
	pos := initialPos

	for _, targetKey := range input {
		targetPos := getPos(grid, targetKey)
		moves := getShortestPath(grid, pos, targetPos)

		for _, move := range moves {
			seq += string(move.dir)
		}

		seq += "A"
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

var (
	directions = []vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	keys       = []rune{'^', '>', 'v', '<'}
)

type vertex struct {
	pos vec
	dir rune
}

func getShortestPath(grid []string, start, end vec) []vertex {
	parent := make(map[vertex]vertex)
	path := make([]vertex, 0)

	visited := make(map[vertex]bool)
	visited[vertex{start, 'A'}] = true

	queue := list.New()
	queue.PushBack(vertex{start, 'A'})

	for queue.Len() > 0 {
		node := queue.Front()
		current := node.Value.(vertex)
		queue.Remove(node)

		if current.pos == end {
			path = make([]vertex, 0)
			for current.pos != start {
				path = append(path, current)
				current = parent[current]
			}

			return path
		}

		for _, neighbor := range getNeighbors(grid, current.pos) {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}

	return []vertex{}
}

func getNeighbors(grid []string, pos vec) []vertex {
	neighbors := make([]vertex, 0)

	for i, dir := range directions {
		neighborPos := vec{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(grid, neighborPos.x, neighborPos.y)

		if ok && neighborId != '_' {
			neighbors = append(neighbors, vertex{neighborPos, keys[i]})
		}
	}

	return neighbors
}
