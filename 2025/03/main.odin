package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"

find_largest :: proc(voltages: []int) -> int {
	number := ""
	index := -1

	for n in 0 ..< 12 {
		r := 0

		for v, i in voltages[:len(voltages) - 12 + 1 + n] {
			if i <= index do continue

			if v > r {
				r = v
				index = i
			}
		}

		number = strings.join([]string{number, fmt.tprintf("%v", r)}, "")
	}

	return strconv.atoi(number)
}

main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))
	sum := 0

	for line in lines {
		voltages := slice.mapper(strings.split(line, ""), proc(b: string) -> int {
			return strconv.atoi(b)
		})

		largest := find_largest(voltages)
		sum += largest
	}

	fmt.println(sum)
}
