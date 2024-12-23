package main

import (
	"fmt"

	"aoc/internal/utils"
)

func main() {
	lines := utils.GetLines("example")

	for _, l := range lines {
		fmt.Println(l)
	}
}
