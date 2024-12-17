package main

import (
	"aoc/internal/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := utils.GetLines("src/2024/04/input")

	// rotated45 := rotate45(lines)
	// rotated90 := rotate90(lines)

	total := 0

	pattern, _ := regexp.Compile(`XMAS`)

	for range 8 {
		for _, line := range lines {
			total += len(pattern.FindAllString(strings.ReplaceAll(line, " ", ""), 999))
		}
		lines = rotate45(lines)
	}

	fmt.Println(total)
}

func rotate45(input []string) []string {
	n := len(input) // Number of rows
	if n == 0 {
		return []string{}
	}

	newSize := 2*n - 1
	newMatrix := make([][]string, newSize)
	for i := range newMatrix {
		newMatrix[i] = make([]string, newSize)
	}

	for i := range newMatrix {
		for j := range newMatrix[i] {
			newMatrix[i][j] = " "
		}
	}

	offset := n - 1
	for i := 0; i < n; i++ {
		for j, char := range input[i] {
			newRow := i + j
			newCol := offset + (j - i)
			newMatrix[newRow][newCol] = string(char)
		}
	}

	output := make([]string, newSize)
	for i := range newMatrix {
		output[i] = strings.Join(newMatrix[i], "")
	}

	return output
}

func rotate90(input []string) []string {
	if len(input) == 0 {
		return input
	}

	rows := len(input)
	cols := len(input[0])

	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-i-1] = rune(input[i][j])
		}
	}

	result := make([]string, cols)
	for i := 0; i < cols; i++ {
		result[i] = string(rotated[i])
	}

	return result
}
