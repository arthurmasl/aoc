package main

import (
	"fmt"
	"slices"

	"aoc/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2020/05/input")
	ids := []int{}

	for _, line := range lines {
		rows := line[:7]
		cols := line[7:]

		row := getRes(127, rows, 'F', 'B')
		col := getRes(7, cols, 'L', 'R')

		id := row*8 + col
		ids = append(ids, id)
		fmt.Println(row, col, id)
	}

	maxId := 0
	for _, id := range ids {
		if id > maxId {
			maxId = id
		}
	}

	slices.Sort(ids)
	fmt.Println(ids)

	for i, id := range ids {
		if id == 45 {
			continue
		}
		if id == 953 {
			break
		}
		if id-1 != ids[i-1] {
			fmt.Println(id)
		}
		if id+1 != ids[i+1] {
			fmt.Println(id)
		}
	}
}

func getRes(init int, rows string, leftSign, rightSign rune) int {
	left := 0
	right := init

	for _, sign := range rows {
		if sign == leftSign {
			right = left + (right-left)/2
		}
		if sign == rightSign {
			left = (left + (right-left)/2) + 1
		}
	}

	return left
}
