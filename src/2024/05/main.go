package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2024/05/input", "\n\n")
	ruleLines := strings.Split(lines[0], "\n")
	pageLines := strings.Split(lines[1], "\n")

	rules := make(map[int][]int)

	for _, ruleLine := range ruleLines {
		keyStr, valueStr, _ := strings.Cut(ruleLine, "|")
		key, _ := strconv.Atoi(keyStr)
		value, _ := strconv.Atoi(valueStr)

		_, ok := rules[key]
		if !ok {
			rules[key] = make([]int, 0)
		}

		rules[key] = append(rules[key], value)
	}

	res := 0

PL:
	for _, pages := range pageLines {
		pageInts := utils.ConvertToInts(strings.Split(pages, ","))

		for pageIndex, page := range pageInts {
			// fmt.Println(page, rules[page])

			for _, right := range rules[page] {
				rightIndex := slices.Index(pageInts, right)
				if rightIndex != -1 && rightIndex < pageIndex {
					// fmt.Println("bad")
					continue PL
				}
			}
		}

		// fmt.Println("correct")
		res += pageInts[len(pageInts)/2]
	}

	fmt.Println("====", res)
}
