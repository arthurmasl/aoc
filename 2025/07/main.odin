package main

import "core:fmt"
import "core:strings"

Vec2 :: [2]int

splits := 0
start_x: int
lines: []string

main :: proc() {
	file := #load("example.txt", string)
	lines = strings.split_lines(strings.trim_right(file, "\n"))

	for i, j in lines[0] {
		if i == 'S' {
			start_x = j
			break
		}
	}


	iter :: proc(pos: Vec2) {
		pos := pos

		for {
			pos.y += 1

			if pos.y + 1 == len(lines) {
				splits += 1
				if splits % 10000 == 0 {
					fmt.println("### count split", splits, pos)
				}
				break
			}

			if lines[pos.y][pos.x] == '^' {
				if lines[pos.y + 1][pos.x - 1] == '.' {
					iter({pos.x - 1, pos.y})
				}
				if lines[pos.y + 1][pos.x + 1] == '.' {
					iter({pos.x + 1, pos.y})
					break
				}
			}

		}
	}

	iter(Vec2{start_x, 0})

	// for line in lines {
	// 	fmt.println(line)
	// }

	fmt.println(splits)
}

// 1764 too high
