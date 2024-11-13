package main

import (
	"fmt"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2020/01/input")
	nums := utils.ConvertToInts(lines)

	for _, numA := range nums {
		for _, numB := range nums {
			for _, numC := range nums {
				if numA+numB+numC == 2020 {
					fmt.Println(numA * numB * numC)
					return
				}
			}
		}
	}
}
