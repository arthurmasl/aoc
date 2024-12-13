package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/src/internal/utils"
)

const add = 10000000000000

func main() {
	lines := utils.GetLines("input", "\n\n")
	result := 0

	for _, block := range lines {
		lines := strings.Split(block, "\n")

		ax, ay := getValues(lines[0], "+")
		bx, by := getValues(lines[1], "+")
		tx, ty := getValues(lines[2], "=")

		tx += add
		ty += add

		d := ax*by - bx*ay
		dx := tx*by - bx*ty
		dy := ax*ty - tx*ay

		foundX := dx == (dx/d)*d
		foundY := dy == (dy/d)*d

		if foundX && foundY {
			result += (dx/d)*3 + (dy / d)
		}

		// possibleX := getPossibleValues(ax, bx, targetX+add)
		// possibleY := getPossibleValues(ay, by, targetY+add)
		// intersections := getIntersections(possibleX, possibleY)
		// if len(intersections) > 0 {
		// 	price := intersections[0][0]*3 + intersections[0][1]
		// 	result += price
		// 	fmt.Println("===", price)
		// }
	}

	fmt.Println(result)
}

func getPossibleValues(coefX, coefY, target int) [][2]int {
	values := make([][2]int, 0)

	for x := 0; coefX*x <= target; x++ {
		for y := 0; coefX*x+coefY*y <= target; y++ {
			if coefX*x+coefY*y == target {
				values = append(values, [2]int{x, y})
			}
		}
	}

	return values
}

func getIntersections(a, b [][2]int) [][2]int {
	intersections := make([][2]int, 0)
	for _, v1 := range a {
		for _, v2 := range b {
			if v1[0] == v2[0] && v1[1] == v2[1] {
				intersections = append(intersections, v1)
			}
		}
	}

	return intersections
}

func getValues(str string, separator string) (int, int) {
	l, r, _ := strings.Cut(str, ", ")
	_, xStr, _ := strings.Cut(l, separator)
	_, yStr, _ := strings.Cut(r, separator)
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)

	return x, y
}
