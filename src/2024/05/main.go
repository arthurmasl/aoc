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
	incorrects := make([][]int, 0)

PL:
	for _, pages := range pageLines {
		pageInts := utils.ConvertToInts(strings.Split(pages, ","))

		for pageIndex, page := range pageInts {
			for _, right := range rules[page] {
				rightIndex := slices.Index(pageInts, right)

				if rightIndex != -1 && rightIndex < pageIndex {
					incorrects = append(incorrects, pageInts)
					continue PL
				}
			}
		}

	}

	for _, pages := range incorrects {
		checkAndFix(pages, rules)
		res += pages[len(pages)/2]
	}

	fmt.Println(res)
}

func checkAndFix(pages []int, rules map[int][]int) {
PL:
	for pageIndex, page := range pages {
		for _, right := range rules[page] {
			rightIndex := slices.Index(pages, right)

			if rightIndex != -1 && rightIndex < pageIndex {
				pages = slices.Delete(pages, pageIndex, pageIndex+1)
				pages = slices.Insert(pages, rightIndex, page)
				checkAndFix(pages, rules)
				continue PL
			}
		}
	}
}
