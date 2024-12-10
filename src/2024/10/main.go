package main

import (
	"aoc/src/internal/utils"
	"fmt"
)

type Node struct {
	x, y int
}

var (
	paths   = 0
	visited = make(map[Node]map[Node]bool)
)

func main() {
	lines := utils.GetLines("src/2024/10/example")

	for y, row := range lines {
		for x, col := range row {
			if col == '0' {
				root := Node{x, y}
				visited[root] = make(map[Node]bool)
				searchNext(lines, root, root)
			}
		}
	}

	fmt.Println(paths)
}

func searchNext(list []string, root, node Node) {
	if visited[root][node] {
		return
	}

	curr := list[node.y][node.x]

	if curr == '9' {
		visited[root][node] = true
		paths++
	}

	safeSet := func(y, x int) {
		if y >= 0 && y < len(list) && x >= 0 && x < len(list[y]) {
			next := list[y][x]
			if curr+1 == next {
				// fmt.Printf("(%v) %v:%v\n", string(next), y, x)
				searchNext(list, root, Node{x, y})
			}
		}
	}

	safeSet(node.y, node.x-1)
	safeSet(node.y, node.x+1)
	safeSet(node.y-1, node.x)
	safeSet(node.y+1, node.x)
}
