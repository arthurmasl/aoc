package main

import (
	"fmt"

	"aoc/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2020/01/input")
	nums := utils.ConvertToInts(lines)

	fmt.Println(solution3(nums))
}

func solution1(nums []int) int {
	for _, a := range nums {
		for _, b := range nums {
			for _, c := range nums {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return -1
}

func solution2(nums []int) int {
	for i, a := range nums {
		for _, b := range nums[i+1:] {
			for _, c := range nums[i+2:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return -1
}

func solution3(nums []int) int {
	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			for c := b + 1; c < len(nums); c++ {
				if nums[a]+nums[b]+nums[c] == 2020 {
					return nums[a] * nums[b] * nums[c]
				}
			}
		}
	}

	return -1
}
