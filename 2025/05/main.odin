package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"


main :: proc() {
	file := #load("input.txt", string)
	chunks := strings.split(file, "\n\n")
	available := u64(0)

	ranges := slice.mapper(
		strings.split_lines(strings.trim_right(chunks[0], "\n")),
		proc(s: string) -> [2]u64 {
			st := strings.split(s, "-")
			n1, _ := strconv.parse_u64_of_base(st[0], 10)
			n2, _ := strconv.parse_u64_of_base(st[1], 10)
			return [2]u64{n1, n2}
		},
	)

	for &range in ranges {
		for &r2 in ranges {
			if range != r2 && range[0] >= r2[0] && range[0] <= r2[1] {
				new_range := [2]u64{range[0], r2[1]}

				if new_range == range || range[0] >= r2[0] && range[1] <= r2[1] {
					new_range = {0, 0}
				} else {
					new_range = {r2[1] + 1, range[1]}
				}

				range = new_range
			}
		}
	}


	ranges_map := map[[2]u64]bool{}
	for range in ranges {
		if range == {0, 0} do continue
		ranges_map[range] = true
	}

	for range in ranges_map {
		available += range[1] - range[0] + 1
	}

	fmt.println(available)
	assert(available == 344306344403172)
}
