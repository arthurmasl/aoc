package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputDir = "src/2020/01/input"

func main() {
	input, _ := os.ReadFile(inputDir)
	lines := strings.Split(string(input)[:len(input)-1], "\n")

	var nums []int

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

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

	fmt.Println(nums)
}
