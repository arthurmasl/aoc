package main

import (
	"fmt"

	"aoc/internal/utils"
)

type Node struct {
	x, y int
}

var (
	paths   = 0
	visited = make(map[Node]map[Node]bool)
)

func main() {
	lines := utils.GetLines("src/2024/10/input")

	for y, row := range lines {
		for x, col := range row {
			if col == '0' {
				root := Node{x, y}
				// visited[root][node] = make(map[Node]bool)
				searchNext(lines, root, root)
			}
		}
	}

	fmt.Println(paths)
}

func searchNext(list []string, root, node Node) {
	// if visited[root][node] {
	// 	return
	// }

	curr := list[node.y][node.x]

	if curr == '9' {
		// visited[root][node] = true
		paths++
	}

	checNext := func(y, x int) {
		if y >= 0 && y < len(list) && x >= 0 && x < len(list[y]) {
			if curr+1 == list[y][x] {
				searchNext(list, root, Node{x, y})
			}
		}
	}

	checNext(node.y, node.x-1)
	checNext(node.y, node.x+1)
	checNext(node.y-1, node.x)
	checNext(node.y+1, node.x)
}
