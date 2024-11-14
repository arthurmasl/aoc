package main

import (
	"fmt"
	"time"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2020/01/example")
	nums := utils.ConvertToInts(lines)

	t1 := time.Now()
	solution2(nums)
	fmt.Println(time.Since(t1))
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
		for i2, b := range nums[i+1:] {
			for _, c := range nums[i2+1:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return -1
}
