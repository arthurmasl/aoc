package main

import "core:fmt"
import "core:strconv"
import "core:strings"


main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	sum: int

	calc :: proc(n: int, sum: ^int) {
		res := n / 3 - 2

		if res >= 0 {
			sum^ += res
			calc(res, sum)
		}
	}

	for line in lines {
		n := strconv.atoi(line)
		calc(n, &sum)
	}

	fmt.println(sum)
}
