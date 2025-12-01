package main

import "core:fmt"
import "core:strconv"
import "core:strings"

main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	size := 100
	index := 50
	wraps := 0

	for line in lines {
		dir := rune(line[0])
		num := strconv.atoi(line[1:])

		wraps += abs(num) / size

		delta := 0
		if dir == 'R' do delta = num % size
		if dir == 'L' do delta = -(num % size)

		if index > 0 && index + delta <= 0 || index + delta >= size {
			wraps += 1
		}

		index = (index + size + delta) % size

	}

	fmt.println(wraps)
}
