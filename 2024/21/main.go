package main

import (
	"container/list"
	"fmt"
	"strconv"
	"unicode/utf8"

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

func main() {
	inputs := utils.GetLines("example")

	total := 0
	for _, input := range inputs[4:5] {
		seq := getSequence(numeric, numericPos, input)
		fmt.Println(seq, len(seq))
		for range 2 {
			seq = getSequence(directional, directionalPos, seq)
			fmt.Println(seq, len(seq))
		}
		fmt.Println()

		fmt.Println(len(seq))
		code, _ := strconv.Atoi(input[:len(input)-1])
		total += code * utf8.RuneCountInString(seq)
	}

	fmt.Println(total)
	utils.Assert(total == 126384)

	// num
	// pos := vec{2, 2}
	// m := getShortstMoves(numeric, vec{0, 0}, numericPos)
	// fmt.Println(m)

	// dir
	// pos := vec{1, 0}
	// m := getShortstMoves(directional, directionalPos, pos)
	// fmt.Println(m)
}

type keys struct {
	from rune
	to   rune
}

func getSequence(grid []string, initialPos vec, input string) string {
	seq := ""
	pos := initialPos
	lastDir := 'A'

	for _, targetKey := range input {
		targetPos := getPos(grid, targetKey)
		moves := getShortstMoves(grid, pos, targetPos, lastDir)
		if len(moves) > 0 {
			lastDir = rune(moves[len(moves)-1])
		}

		seq += moves + "A"
		pos = targetPos
	}

	return seq
}

func getShortstMoves(grid []string, start, end vec, dir rune) string {
	parent := make(map[vec]neighbor)
	path := ""

	visited := make(map[vec]bool)
	visited[start] = true

	queue := list.New()
	queue.PushBack(neighbor{start, dir})

	for queue.Len() > 0 {
		node := queue.Front()
		current := node.Value.(neighbor)
		queue.Remove(node)

		if current.pos == end {
			path = ""
			for current.pos != start {
				path += string(current.dir)
				current = parent[current.pos]
			}
			path = reverseString(path)
			return path
		}

		for _, neighbor := range getNeighbors(grid, current.pos) {
			if !visited[neighbor.pos] && current.dir == neighbor.dir {
				queue.PushBack(neighbor)
				visited[neighbor.pos] = true
				parent[neighbor.pos] = current
				break
			}
		}

		for _, neighbor := range getNeighbors(grid, current.pos) {
			if !visited[neighbor.pos] {
				queue.PushBack(neighbor)
				visited[neighbor.pos] = true
				parent[neighbor.pos] = current
			}
		}
	}

	return path
}

var (
	directions = []vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	arrows     = []rune{'^', '>', 'v', '<'}
)

type neighbor struct {
	pos vec
	dir rune
}

func getNeighbors(grid []string, pos vec) []neighbor {
	neighbors := make([]neighbor, 0)

	for i, dir := range directions {
		neighborPos := vec{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(grid, neighborPos.x, neighborPos.y)

		if ok && neighborId != '_' {
			neighbors = append(neighbors, neighbor{neighborPos, arrows[i]})
		}
	}

	return neighbors
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

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
