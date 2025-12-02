package main

import "core:fmt"
import "core:strconv"
import "core:strings"

main :: proc() {
	file := #load("example.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))
	ranges_strings := strings.split(lines[0], ",")
	result: int

	for range in ranges_strings {
		ranges := strings.split(range, "-")
		from := strconv.atoi(ranges[0])
		to := strconv.atoi(ranges[1])

		fmt.println(from, to)

		for i := from; i <= to; i += 1 {
			s := fmt.tprintf("%v", i)
			if s[0] == 0 || len(s) % 2 != 0 {
				continue
			}

			mid_index := len(s) / 2
			left := s[:mid_index]
			right := s[mid_index:]

			if left == right {
				result += i
			}
		}
	}

	fmt.println(result)
}
