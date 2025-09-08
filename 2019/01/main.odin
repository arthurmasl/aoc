package main

import "core:fmt"
import "core:strconv"
import "core:strings"

main :: proc() {
	file := #load("input", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	sum: int

	for line in lines {
		n := strconv.atoi(line)
		res := n / 3 - 2
		sum += res
	}

	fmt.println(sum)

}
