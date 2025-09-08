package main

import "core:fmt"
import "core:strings"

main :: proc() {
	file := #load("example.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	fmt.println(lines)
}
