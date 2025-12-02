package main

import "core:fmt"
import "core:strconv"
import "core:strings"

Chunk_Iterator :: struct($T: typeid) {
	slice: []T,
	size:  int,
	index: int,
}

chunk_iterator :: proc($T: typeid, it: ^Chunk_Iterator(T)) -> (chunk: []T, ok: bool) {
	if it.index >= len(it.slice) do return nil, false

	end := min(it.index + it.size, len(it.slice))
	chunk = it.slice[it.index:end]

	it.index = end

	return chunk, true
}

main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))
	ranges_strings := strings.split(lines[0], ",")
	result: int

	for range in ranges_strings {
		ranges := strings.split(range, "-")
		from := strconv.atoi(ranges[0])
		to := strconv.atoi(ranges[1])

		for i := from; i <= to; i += 1 {
			s := fmt.tprintf("%v", i)
			if s[0] == 0 {
				continue
			}

			mid_index := len(s) / 2
			ss := strings.split(s, "")

			for j := 1; j <= mid_index; j += 1 {
				it := Chunk_Iterator(string) {
					slice = ss,
					size  = j,
				}

				prev: string
				similar := true
				for val in chunk_iterator(string, &it) {
					p := strings.join(val, "")
					if prev == "" do prev = p
					if prev != p do similar = false
					prev = p
				}

				if similar {
					result += i
					break
				}
			}

		}
	}

	fmt.println(result)
}
