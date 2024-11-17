package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetLines(inputDir string, args ...string) []string {
	input, err := os.ReadFile(inputDir)
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil
	}

	sep := "\n"
	if len(args) > 0 {
		sep = args[0]
	}

	lines := strings.Split(string(input)[:len(input)-1], sep)
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

type mapFunc[E any] func(E) E

func Map[S ~[]E, E any](s S, f mapFunc[E]) S {
	result := make(S, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

type keepFunc[E any] func(E) bool

func Filter[S ~[]E, E any](s S, f keepFunc[E]) S {
	result := S{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

type reduceFunc[E any] func(cur, next E) E

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}
