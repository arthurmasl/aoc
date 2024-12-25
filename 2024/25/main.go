package main

import (
	"fmt"
	"strings"

	"aoc/internal/utils"
)

var (
	filled = "#####"
	empty  = "....."
)

func main() {
	blocks := utils.GetLines("input", "\n\n")

	keys := make([]string, 0)
	locks := make([]string, 0)

	for _, block := range blocks {
		if block[0:5] == filled {
			locks = append(locks, block)
		}
		if block[len(block)-5:] == filled {
			block = strings.ReplaceAll(block, "#", "@")
			keys = append(keys, block)
		}
	}

	pairs := 0
	for _, lock := range locks {
	KL:
		for _, key := range keys {
			for i, k := range lock {
				if k == '#' && key[i] != '.' {
					continue KL
				}
			}
			pairs++
		}
	}

	fmt.Println(pairs)
}
