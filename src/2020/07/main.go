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
	lines := utils.GetLines("src/2020/07/input")
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

	check("shiny gold bag", bags)

	fmt.Println(count)
}

func check(name string, bags Bags) {
	for n, c := range bags[name] {
		for range c {
			count++
			check(n, bags)
		}
	}
}
