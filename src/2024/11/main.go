package main

import (
	"aoc/src/internal/utils"
	"fmt"
	"strings"
	"time"
)

func main() {
	line := utils.GetLines("src/2024/11/input")[0]
	numbers := utils.ConvertToInts(strings.Split(line, " "))

	// stopProfiling := utils.Profile()
	// defer stopProfiling()

	newNumbers := make([]int, 0)

	fmt.Println(numbers)
	t1 := time.Now()

	for iteration := range 25 {
		// i := 0
		for i := range numbers {
			if numbers[i] == 0 {
				// numbers[i] = 1
				// newNumbers[i] = 1
				newNumbers = append(newNumbers, 1)
				// i += 1
				continue
			}

			if digits := countDigits(numbers[i]); digits%2 == 0 {
				l, r := splitNumber(numbers[i], digits/2)
				// numbers = slices.Replace(numbers, i, i+1, l, r)
				// newNumbers[i] = l
				// newNumbers[i+1] = r
				newNumbers = append(newNumbers, l, r)
				// i += 2
				continue
			}

			// numbers[i] = numbers[i] * 2024
			// newNumbers[i] = numbers[i] * 2024
			newNumbers = append(newNumbers, numbers[i]*2024)
			// i += 1
		}

		numbers = make([]int, len(newNumbers))
		copy(numbers, newNumbers)
		newNumbers = nil

		fmt.Println(iteration, len(numbers))
	}

	fmt.Printf("answer: %v time: %v\n", len(numbers), time.Since(t1))
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func splitNumber(num, splitPosition int) (int, int) {
	// Calculate the divisor in a fast way
	divisor := 1
	for i := 0; i < splitPosition; i++ {
		divisor *= 10
	}

	// Use integer division and modulo to split the number
	leftPart := num / divisor
	rightPart := num % divisor

	return leftPart, rightPart
}
