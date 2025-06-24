package main

import "core:fmt"
import "core:os"
import "core:strings"

main :: proc() {
	file, ok := os.read_entire_file("assets/example.txt", context.temp_allocator)
	if !ok do panic("no file")

	lines := strings.split_lines(string(file), context.temp_allocator)
	lines = lines[:len(lines) - 1]

	fmt.println(lines)

	free_all(context.temp_allocator)
}
