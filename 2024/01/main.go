package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2024/01/input")
	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, line := range lines {
		n1, n2, _ := strings.Cut(line, "   ")
		list1[i], _ = strconv.Atoi(n1)
		list2[i], _ = strconv.Atoi(n2)
	}

	fmt.Println(part1(list1, list2))
	fmt.Println(part2(list1, list2))
}

func part1(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	res := 0
	for i, n1 := range list1 {
		n2 := list2[i]
		res += max(n1, n2) - min(n1, n2)
	}

	return res
}

func part2(list1, list2 []int) int {
	res := 0
	for _, n1 := range list1 {
		count := 0
		for _, n2 := range list2 {
			if n1 == n2 {
				count++
			}
		}
		res += n1 * count
	}

	return res
}
