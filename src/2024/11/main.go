package main

import (
	"aoc/src/internal/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	line := utils.GetLines("src/2024/11/input")[0]
	numbers := utils.ConvertToInts(strings.Split(line, " "))

	stopProfiling := utils.Profile()
	defer stopProfiling()

	newNumbers := make([]int, 0)

	fmt.Println(numbers)
	t1 := time.Now()

	for iteration := range 25 {
		for i := range numbers {
			if numbers[i] == 0 {
				newNumbers = append(newNumbers, 1)
				continue
			}

			if digits := countDigits(numbers[i]); digits%2 == 0 {
				l, r := splitNumber(numbers[i], digits/2)
				newNumbers = append(newNumbers, l, r)
				continue
			}

			newNumbers = append(newNumbers, numbers[i]*2024)
		}

		numbers = make([]int, len(newNumbers))
		copy(numbers, newNumbers)
		newNumbers = nil

		fmt.Println(iteration, len(numbers))
	}

	fmt.Printf("answer: %v time: %v\n", len(numbers), time.Since(t1))
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func splitNumber(number, position int) (int, int) {
	divisor := int(math.Pow10(position))

	left := number / divisor
	right := number % divisor

	return left, right
}
