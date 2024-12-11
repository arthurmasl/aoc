package main

import (
	"aoc/src/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	line := utils.GetLines("src/2024/11/input")[0]
	numbers := utils.ConvertToInts(strings.Split(line, " "))

	newNumbers := make([]int, 0)

	for i := range 75 {
		for _, n := range numbers {
			s := strconv.Itoa(n)

			if n == 0 {
				newNumbers = append(newNumbers, 1)
			} else if len(s)%2 == 0 {
				l, r := s[:len(s)/2], s[len(s)/2:]
				lInt, _ := strconv.Atoi(l)
				rInt, _ := strconv.Atoi(r)

				newNumbers = append(newNumbers, lInt, rInt)
			} else {
				newNumbers = append(newNumbers, n*2024)
			}
		}

		numbers = make([]int, len(newNumbers))
		copy(numbers, newNumbers)
		fmt.Println(i+1, len(numbers))
		newNumbers = nil
	}

	// fmt.Println(numbers)
	fmt.Println(len(numbers))
}
