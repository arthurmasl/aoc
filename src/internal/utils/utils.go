package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetLines(inputDir string) []string {
	input, err := os.ReadFile(inputDir)
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil
	}

	lines := strings.Split(string(input)[:len(input)-1], "\n")
	return lines
}

func ConvertToInts(strings []string) []int {
	var nums []int

	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	return nums
}
