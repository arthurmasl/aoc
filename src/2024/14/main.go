package main

import (
	"fmt"
	"strings"

	"aoc/src/internal/utils"
)

type Entity struct {
	pos, vel Vector
}

type Vector struct {
	x, y int
}

type Rect struct {
	x, y, w, h int
}

func main() {
	lines := utils.GetLines("input")
	guards := make([]Entity, len(lines))
	grid := make([]string, 103)

	// draw grid
	for y := range grid {
		grid[y] = strings.Repeat(".", 101)
	}

	// w, h := len(grid[0]), len(grid)
	// halfW := w / 2
	// halfH := h / 2

	// quadrants := []Rect{
	// 	{0, 0, halfW - 1, halfH - 1},
	// 	{halfW + 1, 0, halfW - 1, halfH - 1},
	// 	{0, halfH + 1, halfW - 1, halfH - 1},
	// 	{halfW + 1, halfH + 1, halfW - 1, halfH - 1},
	// }

	// counts := [4]int{}

	// parse guards
	for i, line := range lines {
		var posX, posY, velX, velY int
		fmt.Sscanf(line, "p=%v,%v v=%v,%v", &posX, &posY, &velX, &velY)
		guards[i] = Entity{Vector{posX, posY}, Vector{velX, velY}}
	}

	fullDraw(grid, guards, 8000)

	// reader := bufio.NewReader(os.Stdin)
	// time := 0
	// for {
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	//
	// 	if input == "b" {
	// 		time--
	// 		// reset guards
	// 		for i, line := range lines {
	// 			var posX, posY, velX, velY int
	// 			fmt.Sscanf(line, "p=%v,%v v=%v,%v", &posX, &posY, &velX, &velY)
	// 			guards[i] = Entity{Vector{posX, posY}, Vector{velX, velY}}
	// 		}
	// 		fullDraw(grid, guards, time)
	// 		fmt.Println(time)
	// 		continue
	// 	}
	//
	// 	time++
	// 	drawNext(grid, guards)
	// 	fmt.Println(time)
	// }

	// for _, guard := range guards {
	// 	for qi, r := range quadrants {
	// 		if IsPointInside(r, guard.pos) {
	// 			counts[qi]++
	// 		}
	// 	}
	// }

	// draw(grid, guards)
	// drawQuadrants(grid, quadrants)
	// draw(grid, guards)

	// total := counts[0]
	// for _, v := range counts[1:] {
	// 	total *= v
	// }
	// fmt.Println(counts, total)
}

func fullDraw(grid []string, guards []Entity, time int) {
	t := 0
	for range time {
		t++
		for i, guard := range guards {
			movedPos := Vector{guard.pos.x + guard.vel.x, guard.pos.y + guard.vel.y}
			guards[i].pos = getSafeValue(grid, movedPos)
		}
		draw(grid, guards)

		found := false
		for y, row := range grid {
			for range row {
				if y == 70 {
					if strings.Contains(row, strings.Repeat("O", 30)) {
						found = true
						break
					}
				}
			}
		}

		if found {
			fmt.Println(t)
			for _, row := range grid {
				fmt.Println(row)
			}
		}

	}
}

func drawNext(grid []string, guards []Entity) {
	for i, guard := range guards {
		movedPos := Vector{guard.pos.x + guard.vel.x, guard.pos.y + guard.vel.y}
		guards[i].pos = getSafeValue(grid, movedPos)
	}

	draw(grid, guards)
}

func draw(grid []string, guards []Entity) {
	for y := range grid {
		grid[y] = strings.Repeat(" ", 101)
	}

	for _, guard := range guards {
		newRow := []rune(grid[guard.pos.y])
		if guard.pos.y == 70 {
			newRow[guard.pos.x] = 'O'
		} else {
			newRow[guard.pos.x] = 'X'
		}
		grid[guard.pos.y] = string(newRow)
	}
}

func drawQuadrants(grid []string, quadrants []Rect) {
	for y, row := range grid {
		for x := range row {
			for _, rect := range quadrants {
				if IsPointInside(rect, Vector{x, y}) {
					newRow := []rune(grid[y])
					newRow[x] = 'X'
					grid[y] = string(newRow)
				}
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}

func getSafeValue(arr []string, pos Vector) Vector {
	if pos.y >= 0 && pos.y < len(arr) && pos.x >= 0 && pos.x < len(arr[0]) {
		return pos
	}

	if pos.x < 0 {
		pos.x = len(arr[0]) + pos.x
	}
	if pos.y < 0 {
		pos.y = len(arr) + pos.y
	}

	if pos.x >= len(arr[0]) {
		pos.x = pos.x - len(arr[0])
	}
	if pos.y >= len(arr) {
		pos.y = pos.y - len(arr)
	}

	return pos
}

func IsPointInside(r Rect, p Vector) bool {
	return p.x >= r.x && p.x <= r.x+r.w &&
		p.y >= r.y && p.y <= r.y+r.h
}
