package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2024/02/example")
	res := 0

l:
	for _, line := range lines {
		numbersStr := strings.Split(line, " ")
		numbers := make([]int, len(numbersStr))
		for i := range numbersStr {
			num, _ := strconv.Atoi(numbersStr[i])
			numbers[i] = num
		}

		for i := range numbers[:len(numbers)-1] {
			isSafe := checkSafity(i, numbers)
			if !isSafe {
				continue l
			}
		}

		// safe
		res++
	}

	fmt.Println(res)
}

func checkSafity(i int, numbers []int) bool {
	p := numbers[max(0, i-1)]
	a := numbers[i]
	b := numbers[i+1]

	prevDecreasing := checkIsDecreasing(p, a)
	currDecreasing := checkIsDecreasing(a, b)
	if i == 0 {
		prevDecreasing = currDecreasing
	}

	diff := math.Abs(float64(a - b))

	if (diff < 1 || diff > 3) || prevDecreasing != currDecreasing {
		// unsafe
		return false
	}

	// safe
	return true
}

func checkIsDecreasing(a, b int) bool {
	return a >= b
}