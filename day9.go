package main

import (
	"fmt"
	"sort"
)

func main() {
	lines := readFileByLine("day9.txt")
	floor := padArray(parseInput(lines))
	baseins := make(map[int]int)
	for i := 1; i+1 < len(floor); i++ {
		for k := 1; k+1 < len(floor[i]); k++ {
			if floor[i][k] != 9 && floor[i][k] != 10 {
				x, y := findBasin(floor, i, k)
				baseins[x+100*y] += 1
			}
		}
	}
	var sizes []int
	for _, size := range baseins {
		sizes = append(sizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

func findBasin(floor [][]int, y int, x int) (int, int) {
	target := floor[y][x]
	if target > floor[y][x-1] {
		return findBasin(floor, y, x-1)
	} else if target > floor[y][x+1] {
		return findBasin(floor, y, x+1)
	} else if target > floor[y-1][x] {
		return findBasin(floor, y-1, x)
	} else if target > floor[y+1][x] {
		return findBasin(floor, y+1, x)
	} else {
		return x, y
	}
}

func isLowest(floor [][]int, y int, x int) bool {
	target := floor[y][x]
	return target < floor[y][x-1] &&
		target < floor[y][x+1] &&
		target < floor[y-1][x] &&
		target < floor[y+1][x]
}

func padArray(arr [][]int) [][]int {
	for i := range arr {
		arr[i] = append([]int{10}, arr[i]...)
		arr[i] = append(arr[i], 10)
	}

	lineLength := len(arr[0])
	var padLine []int
	for i := 0; i < lineLength; i++ {
		padLine = append(padLine, 10)
	}
	arr = append([][]int{padLine}, arr...)
	arr = append(arr, padLine)
	return arr
}

func parseInput(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		var resultLine []int
		for _, char := range line {
			resultLine = append(resultLine, charToInt(char))
		}
		result = append(result, resultLine)
	}
	return result
}
