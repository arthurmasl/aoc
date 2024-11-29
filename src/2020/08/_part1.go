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
	lines := utils.GetLines("src/2020/08/example")
	run(lines)
}

func run(lines []string) {
	ops := make(map[int]bool)
	acc := 0
	i := 0

	for {
		line := lines[i]
		_, ok := ops[i]
		if !ok {
			ops[i] = true
		} else {
			fmt.Println(line)
			break
		}

		op, offsetStr, _ := strings.Cut(line, " ")
		offset, _ := strconv.Atoi(offsetStr)

		fmt.Println(i, op, offset)
		switch op {
		case "nop":
			i++
			continue
		case "acc":
			acc += offset
			i++
		case "jmp":
			i += offset
		}
	}

	fmt.Println(acc)
}
