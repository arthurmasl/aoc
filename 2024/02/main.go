package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2024/02/input")
	res := 0

lineLoop:
	for _, line := range lines {
		numbersStr := strings.Split(line, " ")
		numbers := make([]int, len(numbersStr))

		for i := range numbersStr {
			num, _ := strconv.Atoi(numbersStr[i])
			numbers[i] = num
		}

		if isSafe := fullCheck(numbers); isSafe {
			res++
			continue lineLoop
		} else {
			for i := range numbers {
				withoutIndex := append(append([]int{}, numbers[:i]...), numbers[i+1:]...)

				isSafe = fullCheck(withoutIndex)
				if isSafe {
					res++
					continue lineLoop
				}
			}
		}

	}

	fmt.Println(res)
}

func fullCheck(numbers []int) bool {
	for i := range numbers[:len(numbers)-1] {
		isSafe := checkSafity(i, numbers)

		if !isSafe {
			return false
		}
	}

	return true
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
