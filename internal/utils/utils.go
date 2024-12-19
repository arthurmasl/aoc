package utils

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

func Assert(condition bool) {
	if !condition {
		panic("Assertion failed")
	}
}

func GetLines(inputDir string, args ...string) []string {
	input, err := os.ReadFile("assets/" + inputDir + ".txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil
	}

	sep := "\n"
	if len(args) > 0 {
		sep = args[0]
	}

	lines := strings.Split(strings.TrimSpace(string(input)), sep)
	return lines
}

func GetSafeValue(arr []string, x, y int) (byte, bool) {
	if y >= 0 && y < len(arr) && x >= 0 && x < len(arr[y]) {
		return arr[y][x], true
	}

	return 0, false
}

func ConvertToInts(strings []string) []int {
	nums := make([]int, len(strings))

	for i, str := range strings {
		num, _ := strconv.Atoi(str)
		nums[i] = num
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

func Window[Slice ~[]E, E any](slice Slice, size int) iter.Seq[Slice] {
	return func(yield func(Slice) bool) {
		for i := range slice[:len(slice)-size+1] {
			if !yield(slice[i : i+size]) {
				return
			}
		}
	}
}

func WindowString(str string, size int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := range str[:len(str)-size+1] {
			if !yield(str[i : i+size]) {
				return
			}
		}
	}
}
