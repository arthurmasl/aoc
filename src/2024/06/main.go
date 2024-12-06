package main

import (
	"fmt"

	"aoc/src/internal/utils"
)

type Vector struct {
	x, y int
}

// 758 to low
// 2372 too high
func main() {
	lines := utils.GetLines("src/2024/06/input")
	w, h := len(lines[0]), len(lines)

	visited := make(map[Vector]bool)

	pos := Vector{}
	startPos := Vector{}
	vel := Vector{0, -1}

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '^' {
				pos = Vector{x: x, y: y}
				startPos = Vector{x: x, y: y}
				visited[pos] = true
			}
		}
	}

	for {
		next := Vector{x: pos.x + vel.x, y: pos.y + vel.y}

		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			break
		}

		if lines[next.y][next.x] == '#' {
			vel.x, vel.y = -vel.y, vel.x
			continue
		}

		pos = next

		// draw
		// newLine := []rune(lines[pos.y])
		// newLine[pos.x] = 'X'
		// lines[pos.y] = string(newLine)

		visited[pos] = true
	}

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

		found := scan(newLines, startPos)
		if !found {
			loops++
		}
	}

	fmt.Println(len(visited), loops)
}

func scan(lines []string, pos Vector) bool {
	w, h := len(lines[0]), len(lines)
	vel := Vector{0, -1}

	tries := 0
	for {
		tries++
		if tries == 10000 {
			return false
		}

		next := Vector{x: pos.x + vel.x, y: pos.y + vel.y}

		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h {
			return true
		}

		if lines[next.y][next.x] == '#' {
			vel.x, vel.y = -vel.y, vel.x
			continue
		}

		pos = next
	}
}
