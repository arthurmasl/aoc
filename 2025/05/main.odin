package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"


main :: proc() {
	file := #load("example.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	for line in lines {
		fmt.println(line)
	}

}
