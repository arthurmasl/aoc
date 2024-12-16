package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"

	"aoc/src/internal/utils"
)

type Vector struct {
	x, y float64
}

// t r b l
var directions = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

// 453 to low
func main() {
	lines := utils.GetLines("example")

	startPos := Vector{}
	endPos := Vector{}

	distance := make(map[Vector]int)
	parent := make(map[Vector]Vector)
	path := make([]Vector, 0)
	lowestCost := 0

	paths := make(map[int][][]Vector)

	for y, row := range lines {
		for x, col := range row {
			if col == 'S' {
				startPos = Vector{float64(x), float64(y)}
			}
			if col == 'E' {
				endPos = Vector{float64(x), float64(y)}
			}
		}
	}

	priorityQueue := make(PriorityQueue, 0)
	heap.Push(&priorityQueue, &Item{value: ItemValue{startPos, directions[1], 0}})

	distance[startPos] = 0

	for priorityQueue.Len() > 0 {
		element := heap.Pop(&priorityQueue).(*Item)
		currentPos := element.value.pos
		currentDir := element.value.direction
		currentCost := element.value.cost

		if currentPos == endPos {
			path = make([]Vector, 0)
			for currentPos != startPos {
				path = append(path, currentPos)
				currentPos = parent[currentPos]
			}
			slices.Reverse(path)
			lowestCost = currentCost
			paths[currentCost] = append(paths[currentCost], path)
		}

		for _, neighbor := range getNeighbors(lines, currentPos) {
			neigborDir := GetDirection(currentPos, neighbor)
			newCost := currentCost + 1
			if neigborDir != currentDir {
				newCost += 1000
			}

			if cost, ok := distance[neighbor]; !ok || newCost <= cost {
				distance[neighbor] = newCost
				parent[neighbor] = currentPos
				heap.Push(
					&priorityQueue,
					&Item{value: ItemValue{neighbor, neigborDir, newCost}},
				)
			}
		}
	}

	for _, path := range paths[lowestCost] {
		// lines := utils.GetLines("example")
		// fill path
		for _, pos := range path {
			for y, row := range lines {
				newRow := []byte(row)
				for x := range row {
					if int(pos.x) == x && int(pos.y) == y {
						newRow[x] = 'O'
					}
				}
				lines[y] = string(newRow)
			}
		}
	}

	// draw grid
	for _, row := range lines {
		fmt.Println(row)
	}
	fmt.Println()

	total := 1
	for _, row := range lines {
		for _, col := range row {
			if col == 'O' {
				total++
			}
		}
	}

	fmt.Println(len(paths[lowestCost]))
	fmt.Println(lowestCost, total)
}

func getNeighbors(lines []string, pos Vector) []Vector {
	neighbors := make([]Vector, 0)

	for _, dir := range directions {
		neighborPos := Vector{pos.x + dir.x, pos.y + dir.y}
		neighborId, ok := utils.GetSafeValue(lines, int(neighborPos.x), int(neighborPos.y))

		if ok && neighborId != '#' {
			neighbors = append(neighbors, neighborPos)
		}
	}

	return neighbors
}

type ItemValue struct {
	pos       Vector
	direction Vector
	cost      int
}

type Item struct {
	value    ItemValue
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value ItemValue, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (v Vector) Normalize() Vector {
	mag := math.Sqrt(v.x*v.x + v.y*v.y)
	if mag == 0 {
		return Vector{0, 0}
	}
	return Vector{x: v.x / mag, y: v.y / mag}
}

func Direction(v1, v2 Vector) Vector {
	dx := v2.x - v1.x
	dy := v2.y - v1.y
	return Vector{x: dx, y: dy}.Normalize()
}

func CardinalDirection(input Vector) Vector {
	normalized := input.Normalize()
	if math.Abs(normalized.x) > math.Abs(normalized.y) {
		if normalized.x > 0 {
			return Vector{1, 0} // Right
		}
		return Vector{-1, 0} // Left
	} else {
		if normalized.y > 0 {
			return Vector{0, 1} // Up
		}
		return Vector{0, -1} // Down
	}
}

func GetDirection(v1, v2 Vector) Vector {
	dir := Direction(v1, v2)
	return CardinalDirection(dir)
}
