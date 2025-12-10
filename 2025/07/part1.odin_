package main

import "core:fmt"
import "core:strings"
import "core:unicode/utf8"


main :: proc() {
	file := #load("example.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	splits := 0

	for prev, li in lines[:len(lines) - 1] {
		next := lines[li + 1]
		new_next := utf8.string_to_runes(next)

		for p, pi in prev {
			if p == 'S' && next[pi] == '.' {
				new_next[pi] = '|'
			}

			if p == '|' && next[pi] == '^' {
				splits += 1
				if next[pi - 1] == '.' {
					new_next[pi - 1] = '|'
				}
				if next[pi + 1] == '.' {
					new_next[pi + 1] = '|'
				}
			}

			if p == '|' && next[pi] == '.' {
				new_next[pi] = '|'
			}
		}

		lines[li + 1] = utf8.runes_to_string(new_next)
	}


	for line in lines {
		fmt.println(line)
	}

	fmt.println(splits)
}

// 1764 too high
