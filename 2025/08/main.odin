package main

import "core:fmt"
import "core:math/linalg"
import "core:sort"
import "core:strconv"
import "core:strings"

Vec3 :: [3]f32

Box :: struct {
	pos:   Vec3,
	group: int,
}

main :: proc() {
	file := #load("example.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))
	boxes := make([dynamic]Box)

	for line in lines {
		d := strings.split(line, ",")
		x, _ := strconv.parse_f32(d[0])
		y, _ := strconv.parse_f32(d[1])
		z, _ := strconv.parse_f32(d[2])
		pos := Vec3{x, y, z}
		append(&boxes, Box{pos = pos})
	}

	// sort.bubble_sort_proc(boxes[:], proc(a, b: Box) -> int {
	// 	return int(linalg.length(a.pos) - linalg.length(b.pos))
	// })


	group_index := 0
	for &box in boxes {
		dist := f32(max(f32))
		closest: ^Box

		for &other in boxes {
			if other == box do continue

			new_distance := linalg.distance(box.pos, other.pos)

			if new_distance < dist {
				dist = new_distance
				closest = &other
			}
		}

		fmt.println(box.pos, closest.pos)

		if box.group == 0 && closest.group == 0 {
			group_index += 1
			box.group = group_index
			closest.group = group_index
			continue
		}

		if box.group != 0 {
			closest.group = box.group
			fmt.println("to group", box.group)
			continue
		}

		if closest.group != 0 {
			box.group = closest.group
			fmt.println("to group", box.group)
			continue
		}

	}

	groups := make(map[int]int)
	for box in boxes {
		groups[box.group] += 1
	}

	result := 1
	for group, count in groups {
		fmt.println(group, count)
		result *= count
	}

	fmt.println(result)
}
