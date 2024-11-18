package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

type Bags map[string]map[string]int

var (
	target = "shiny gold bag"
	count  = 0
)

func main() {
	lines := utils.GetLines("src/2020/07/example")
	bags := make(Bags)

	for _, line := range lines {
		l := strings.ReplaceAll(strings.ReplaceAll(line, "bags", "bag"), ".", "")
		bag := strings.Split(l, " contain")[0]
		contains := strings.Split(strings.Split(l, "contain ")[1], ", ")
		bags[bag] = make(map[string]int)

		for _, c := range contains {
			count, _ := strconv.Atoi(c[:1])
			name := c[2:]
			if !strings.Contains(name, "other bag") {
				bags[bag][name] = count
			}
		}

	}

	for bag, contains := range bags {
		if bag == target {
			continue
		}

		for name := range contains {
			count += goDeeper(name, bags)
		}
	}

	fmt.Println(count)
}

func goDeeper(name string, bags Bags) int {
	if name == target {
		return 1
	}

	for n := range bags[name] {
		if n == target {
			return 1
		}

		if goDeeper(n, bags) == 1 {
			return 1
		}
	}

	return 0
}
