package main

import (
	"fmt"

	"aoc/internal/utils"
)

type Pos struct {
	x, y int
}

var (
	directions = []Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited    = make(map[Pos]bool)
)

func main() {
	lines := utils.GetLines("example")
	canvas := createCanvas(lines)

	for y, row := range lines {
		for x, col := range row {
			pos := Pos{x, y}
			if visited[pos] {
				continue
			}

			currId := string(col)
			queue := list.New()
			queue.PushFront(pos)

			for queue.Len() > 0 {
				curr := queue.Front()
				queue.Remove(curr)
				currPos := curr.Value.(Pos)

				visited[currPos] = true
				draw(canvas)

				for _, dir := range directions {
					nextPos := Pos{currPos.x + dir.x, currPos.y + dir.y}
					if visited[nextPos] {
						continue
					}

					if nextId, ok := getSafeValue(lines, nextPos); ok {
						if nextId == currId {
							queue.PushFront(nextPos)
						}
					}
				}
			}

		}
	}
}

func getSafeValue(arr []string, pos Pos) (string, bool) {
	if pos.y >= 0 && pos.y < len(arr) && pos.x >= 0 && pos.x < len(arr[0]) {
		return string(arr[pos.y][pos.x]), true
	}
	return "", false
}

func createCanvas(lines []string) *[]string {
	return &lines
}

var reset = "\033[0m"

func draw(canvas *[]string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println()
	for y, row := range *canvas {
		for x, char := range row {
			if visited[Pos{x, y}] {
				color := getColorByLetter(char)
				fmt.Print(color + string(char) + reset)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}

	time.Sleep(time.Millisecond * 10)
}

func getColorByLetter(letter rune) string {
	colorCode := (int(letter) - int('A')) + 16
	return fmt.Sprintf("\033[38;5;%dm", colorCode%256)
}
