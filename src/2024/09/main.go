package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"aoc/src/internal/utils"
)

type Item struct {
	id int
}

func main() {
	input, _ := os.ReadFile("src/2024/09/example")

	storage := make([]Item, 0)
	id := 0

	for chunk := range slices.Chunk(input[:len(input)-1], 2) {
		size, _ := strconv.Atoi(string(chunk[0]))
		space := 0
		if len(chunk) == 2 {
			space, _ = strconv.Atoi(string(chunk[1]))
		}
		for range size {
			storage = append(storage, Item{id})
		}
		for range space {
			storage = append(storage, Item{-1})
		}
		id++
	}

	for _, item := range slices.Backward(storage) {
		if item.id == -1 {
			storage = slices.Delete(storage, len(storage)-1, len(storage))
		} else {
			break
		}
	}

	// draw
	for _, n := range storage {
		if n.id == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(n.id)
		}
	}
	fmt.Println()

	for i, cell := range storage {
		if cell.id == -1 {
			for j, n := range slices.Backward(storage) {
				if n.id != -1 {
					storage[i].id = n.id
					storage = slices.Delete(storage, j, j+1)
					break
				}
			}

			// draw
			for _, n := range storage {
				if n.id == -1 {
					fmt.Print(".")
				} else {
					fmt.Print(n.id)
				}
			}
			fmt.Println()
		}
	}

	storage = utils.Filter(storage, func(v Item) bool {
		return v.id != -1
	})
	sum := 0
	for i, item := range storage {
		if item.id == -1 {
			continue
		}
		fmt.Println(i, id)
		sum += i * item.id
	}

	fmt.Println(sum)
}
