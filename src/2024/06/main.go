package main

import (
	"fmt"

	"aoc/src/internal/utils"
)

type Vector struct {
	x, y int
}

// 4987 to low
func main() {
	lines := utils.GetLines("src/2024/06/example")
	w, h := len(lines[0]), len(lines)

	visited := make(map[Vector]bool)

	pos := Vector{}
	vel := Vector{0, -1}

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '^' {
				pos = Vector{x: x, y: y}
				visited[pos] = true
			}
		}
	}

	for {
		next := Vector{x: pos.x + vel.x, y: pos.y + vel.y}

		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			fmt.Println("break")
			break
		}

		if lines[next.y][next.x] == '#' {
			vel.x, vel.y = -vel.y, vel.x
			continue
		}

		pos = next

		// draw
		newLine := []rune(lines[pos.y])
		newLine[pos.x] = 'X'
		lines[pos.y] = string(newLine)

		visited[pos] = true
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println(len(visited))
}
