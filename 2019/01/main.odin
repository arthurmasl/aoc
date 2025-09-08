package main

import "core:fmt"
import "core:strconv"
import "core:strings"


sum: int

main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	calc :: proc(n: int) {
		res := n / 3 - 2

		if res >= 0 {
			sum += res
			calc(res)
		}
	}

	for line in lines {
		n := strconv.atoi(line)
		calc(n)
	}

	fmt.println(sum)
}
