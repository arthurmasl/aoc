package main

import (
	"fmt"

	"aoc/src/internal/utils"
)

type Vector struct {
	x, y int
}

func main() {
	lines := utils.GetLines("src/2024/08/example")
	antennas := make(map[string][]Vector)
	antinodes := make(map[Vector]bool)

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				antennas[string(char)] = append(antennas[string(char)], Vector{x, y})
			}
		}
	}

	for id, list := range antennas {
		fmt.Println(id, list)

		for i, antenna1 := range list {
			for _, antenna2 := range list[i+1:] {
				fmt.Println("antenas", antenna1, antenna2)
				l, r := getDirectedX(antenna1, antenna2)
				t, b := getDirectedY(antenna1, antenna2)
				distanceX := r.x - l.x
				distanceY := t.y - b.y

				var left, right Vector
				if l.x == t.x && l.y == t.y || r.x == b.x && r.y == b.y {
					left = Vector{l.x - distanceX, l.y + distanceY}
					right = Vector{r.x + distanceX, r.y - distanceY}
				}
				if l.x == b.x && l.y == b.y || r.x == t.x && r.y == t.y {
					left = Vector{l.x - distanceX, l.y - distanceY}
					right = Vector{r.x + distanceX, r.y + distanceY}
				}

				fmt.Println()

				safeSet(lines, &antinodes, &left)
				safeSet(lines, &antinodes, &right)
			}
		}
	}

	fmt.Println(len(antinodes))

	for pos := range antinodes {
		newLine := []rune(lines[pos.y])
		newLine[pos.x] = '#'
		lines[pos.y] = string(newLine)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func safeSet(arr []string, target *map[Vector]bool, vector *Vector) {
	w, h := len(arr[0]), len(arr)
	if vector.y >= 0 && vector.y < h && vector.x >= 0 && vector.x < w {
		(*target)[*vector] = true
	}
}

func getDirectedX(a, b Vector) (Vector, Vector) {
	var left, right Vector
	if a.x < b.x {
		left = a
		right = b
	} else {
		left = b
		right = a
	}
	return left, right
}

func getDirectedY(a, b Vector) (Vector, Vector) {
	var top, bottom Vector
	if a.y > b.y {
		top = a
		bottom = b
	} else {
		top = b
		bottom = a
	}
	return top, bottom
}
