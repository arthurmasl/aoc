package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"


main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	h := len(lines) - 1
	w := len(slice.filter(strings.split(lines[0], " "), proc(s: string) -> bool {
			return s != ""
		}))
	fmt.println(h, w)

	numbers := make([][]int, w)
	for &n in numbers {
		n = make([]int, h)
	}
	procs := make([]string, w)

	for line, x in lines {
		s := slice.filter(strings.split(line, " "), proc(s: string) -> bool {
			return s != ""
		})

		for n, y in s {
			if n == "+" || n == "*" {
				procs[y] = n
				continue
			}

			number := strconv.atoi(n)
			numbers[y][x] = number
		}
	}

	fmt.println(numbers)
	fmt.println(procs)

	total := 0
	for g, i in numbers {
		sum := 0
		for n, j in g {
			fmt.println(n, procs[i])
			if procs[i] == "+" {
				sum += n
			}

			if procs[i] == "*" {
				if j == 0 {
					sum += n
				} else {
					sum *= n
				}
			}
		}

		fmt.println("===", sum)
		total += sum
	}

	fmt.println(total)
}
