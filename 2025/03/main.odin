package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"

find_largest :: proc(voltages: []int) -> int {
	left := 0
	index := -1
	for v, i in voltages[:len(voltages) - 12] {
		if v > left {
			left = v
			index = i
		}
	}

	right := ""
	for n in 0 ..< 11 {
		r := 0

		for v, i in voltages[:len(voltages) - 12 + n] {
			if i <= index do continue

			if v > r {
				r = v
				index = i
			}
		}

		right = strings.join([]string{right, fmt.tprintf("%v", r)}, "")
	}


	fmt.println(left, right)
	return strconv.atoi(fmt.tprintf("%v%v", left, right))
}

main :: proc() {
	file := #load("example.txt", string)
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
	assert(sum == 3121910778619)
}
