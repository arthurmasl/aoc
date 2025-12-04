package main

import "core:fmt"
import "core:strings"
import "core:unicode/utf8"

Vec2 :: [2]int

directions := []Vec2{{0, -1}, {1, 0}, {0, 1}, {-1, 0}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}}

main :: proc() {
	file := #load("example.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))
	grid := make([][]rune, len(lines))
	result_grid := make([][]rune, len(lines))

	for row, y in lines {
		grid[y] = utf8.string_to_runes(row)
		result_grid[y] = utf8.string_to_runes(row)
	}

	height := len(grid)
	width := len(grid[0])

	rolls := 0


	for row, y in grid {
		col_loop: for &col, x in row {
			if col == '@' {
				pos := Vec2{x, y}

				count := 0
				dir_loop: for dir in directions {
					// for _ in 0 ..< 2 {
					new_pos := pos + dir

					if new_pos.y < 0 ||
					   new_pos.y >= height ||
					   new_pos.x < 0 ||
					   new_pos.x >= width {
						continue dir_loop
					}


					if grid[new_pos.y][new_pos.x] == '@' {
						count += 1
					}

					if count >= 4 {
						continue col_loop
					}
					// }
				}

				if count < 4 {
					rolls += 1
					result_grid[y][x] = 'x'
				}

			}
		}
	}


	// pos := Vec2{0, 0}
	// pos += Vec2{1, 0}
	// pos += Vec2{1, 0}
	// fmt.println(rune(grid[pos.y][pos.x]))

	for row in result_grid {
		for col in row {
			fmt.print(col)
		}
		fmt.println()
	}

	fmt.println(rolls)

	// str := "abc"
	// new_str := utf8.string_to_runes(str)
	// new_str[0] = 'B'
	// str = utf8.runes_to_string(new_str)
	// fmt.println(new_str)
}
