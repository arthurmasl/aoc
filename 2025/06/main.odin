package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"
import "core:unicode/utf8"


main :: proc() {
	file := #load("input.txt", string)
	lines := strings.split_lines(strings.trim_right(file, "\n"))

	h := len(lines) - 1
	w := len(slice.filter(strings.split(lines[0], " "), proc(s: string) -> bool {
			return s != ""
		}))
	fmt.println(h, w)

	numbers := make([][]string, w)
	for &n in numbers {
		n = make([]string, h)
	}
	procs := make([]string, w)

	for line in lines {
		s := slice.filter(strings.split(line, " "), proc(s: string) -> bool {
			return s != ""
		})

		for n, y in s {
			if n == "+" || n == "*" {
				procs[y] = n
				continue
			}

			// number := n
			// numbers[y][x] = number
		}
	}

	ps := 0
	ni := -1
	ml := len(lines[len(lines) - 1])
	for p, i in lines[len(lines) - 1] {
		if (p == '*' || p == '+') {
			sep := i

			for l, jj in lines[:len(lines) - 1] {
				if ni == -1 do continue
				numbers[ni][jj] = l[ps:sep - 1]
			}
			ps = sep

			ni += 1
		}

		if i == ml - 1 {
			for l, jj in lines[:len(lines) - 1] {
				if ni == -1 do continue
				numbers[ni][jj] = l[ps:]
			}
		}
	}

	fmt.println(numbers)
	fmt.println(procs)


	total := 0
	#reverse for g, gi in numbers {
		l := 0
		p := procs[gi]
		for n in g {
			l = max(len(n), l)
		}

		fmt.println()
		fmt.println("====ml", g, l)
		sum := 0

		for j in 0 ..< l {
			i := l - j - 1

			fn := ""
			#reverse for n in g {
				if len(n) > i {
					// number := strconv.atoi(utf8.runes_to_string([]rune{rune(n[i])}))
					sn := utf8.runes_to_string([]rune{rune(n[i])})
					fn = strings.join([]string{sn, fn}, "")
				}
			}

			wos, _ := strings.replace_all(fn, " ", "")
			number := strconv.atoi(wos)
			fmt.println("fn", p, number)

			if p == "+" {
				sum += number
			}

			if p == "*" {
				if sum == 0 {
					sum += number
				} else {
					sum *= number
				}
			}
		}

		fmt.println("SUM: ", sum)
		total += sum
	}

	fmt.println(total)
	assert(total == 3263827)
}
