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
	lines := utils.GetLines("src/2020/08/input")

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		op := strings.Split(line, " ")[0]
		if op == "acc" {
			continue
		}

		var newLine string
		if op == "jmp" {
			newLine = strings.Replace(line, "jmp", "nop", 1)
		}
		if op == "nop" {
			newLine = strings.Replace(line, "nop", "jmp", 1)
		}

		newLines := make([]string, len(lines))
		copy(newLines, lines)
		newLines[i] = newLine

		acc, terminated := run(newLines)
		if terminated {
			fmt.Println(acc, terminated)
			break
		}
	}
}

func run(lines []string) (int, bool) {
	terminated := false
	ops := make(map[int]bool)
	acc := 0
	i := 0

	for i < len(lines) {
		line := lines[i]
		_, ok := ops[i]
		if !ok {
			ops[i] = true
		} else {
			break
		}

		op, offsetStr, _ := strings.Cut(line, " ")
		offset, _ := strconv.Atoi(offsetStr)

		switch op {
		case "acc":
			acc += offset
			i++
		case "jmp":
			i += offset
		case "nop":
			i++
		}
	}

	if i >= len(lines) {
		terminated = true
	}

	return acc, terminated
}
