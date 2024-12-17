package main

import (
	"fmt"

	"aoc/internal/utils"
)

type Vector struct {
	x, y int
}

func main() {
	lines := utils.GetLines("src/2024/06/example")
	startPos := Vector{}

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '^' {
				startPos = Vector{x: x, y: y}
			}
		}
	}

	_, visited := scan(lines, startPos)
	visited[startPos] = true

	for _, line := range lines {
		fmt.Println(line)
	}

	loops := 0
	for v := range visited {
		if v == startPos {
			continue
		}

		newLine := []rune(lines[v.y])
		newLine[v.x] = '#'

		newLines := make([]string, len(lines))
		copy(newLines, lines)
		newLines[v.y] = string(newLine)

		found, _ := scan(newLines, startPos)
		if !found {
			loops++
		}
	}

	fmt.Println(len(visited), loops)
}

func scan(lines []string, pos Vector) (bool, map[Vector]bool) {
	w, h := len(lines[0]), len(lines)
	vel := Vector{0, -1}
	visited := make(map[Vector]bool)

	tries := 0
	for {
		tries++
		if tries == 10000 {
			return false, visited
		}

		next := Vector{x: pos.x + vel.x, y: pos.y + vel.y}

		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			return true, visited
		}

		if lines[next.y][next.x] == '#' {
			vel.x, vel.y = -vel.y, vel.x
			continue
		}

		pos = next
		visited[pos] = true
		// draw
		// newLine := []rune(lines[pos.y])
		// newLine[pos.x] = 'X'
		// lines[pos.y] = string(newLine)
	}
}
