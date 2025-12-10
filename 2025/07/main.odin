package main

import "core:fmt"
import "core:strings"
import "core:unicode/utf8"

splits := 0

main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))


	iter :: proc(olines: []string, start: int = 0) {
		lines := make([]string, len(olines))
		copy(lines, olines)
		// fmt.println("new iter", start)

		for prev, li in lines[:len(lines) - 1] {
			if li < start do continue

			next := lines[li + 1]

			new_next := utf8.string_to_runes(next)

			for p, pi in prev {
				if li == 0 && p == 'S' && next[pi] == '.' {
					// new_next := utf8.string_to_runes(next)
					new_next[pi] = '|'
				}

				if p == '|' && next[pi] == '.' {
					// new_next := utf8.string_to_runes(next)
					new_next[pi] = '|'
				}
				lines[li + 1] = utf8.runes_to_string(new_next)

				if p == '|' && next[pi] == '^' {
					if next[pi - 1] == '.' {
						nn := utf8.string_to_runes(next)
						nn[pi - 1] = '|'

						// new_lines := make([]string, len(lines))
						// copy(new_lines, lines)
						lines[li + 1] = utf8.runes_to_string(nn)
						iter(lines, li + 1)
					}

					if next[pi + 1] == '.' {
						nn := utf8.string_to_runes(next)
						nn[pi + 1] = '|'

						// new_lines := make([]string, len(lines))
						// copy(new_lines, lines)
						lines[li + 1] = utf8.runes_to_string(nn)
						iter(lines, li + 1)
					}

				}

				if li == len(lines) - 2 && p == '|' {
					splits += 1
					if splits % 1000 == 0 {
						fmt.println("### count split", splits)
					}
					// for line, i in lines {
					// 	fmt.println(i, line)
					// }
				}
			}


		}

	}

	iter(lines)

	// for line in lines {
	// 	fmt.println(line)
	// }

	fmt.println(splits)
}

// 1764 too high
