package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

func main() {
	lines := utils.GetLines("src/2024/07/example")
	res := 0

LL:
	for _, line := range lines {
		expectedStr, numbersStr, _ := strings.Cut(line, ": ")
		expected, _ := strconv.Atoi(expectedStr)
		numbers := utils.ConvertToInts(strings.Split(numbersStr, " "))
		permutaitons := getPermutations([]string{"*", "+", "||"}, len(numbers)-1)

	PL:
		for _, permutation := range permutaitons {
			sum := numbers[0]
			for i, number := range numbers[1:] {
				sign := permutation[i]
				switch sign {
				case "*":
					sum *= number
				case "+":
					sum += number
				case "||":
					sum, _ = strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(number))
				}

				if sum == expected {
					res += sum
					continue LL
				}

				if sum > expected {
					continue PL
				}
			}
		}
	}

	fmt.Println(res)
}

func getPermutations(elements []string, k int) [][]string {
	result := [][]string{}
	permutation := make([]string, k)

	var generate func(depth int)
	generate = func(depth int) {
		if depth == k {
			permutationCopy := make([]string, k)
			copy(permutationCopy, permutation)
			result = append(result, permutationCopy)
			return
		}

		for _, element := range elements {
			permutation[depth] = element
			generate(depth + 1)
		}
	}

	generate(0)
	return result
}
