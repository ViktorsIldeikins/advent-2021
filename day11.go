package main

import "fmt"

func main() {
	board := parseInputDay11(readFileByLine("day11.txt"))
	totalFlashes := 0
	var tempLighted int
	count := 0
	for ; !isAllFlashed(board); count++ {
		board, tempLighted = doStep(board)
		totalFlashes += tempLighted
	}
	fmt.Println(count)
}

func doStep(board [][]int) ([][]int, int) {
	board = increment(board)
	var allLighted []IntPair
	currentLighted := findLighted(board)
	for len(allLighted) != len(currentLighted) {
		unique := findUnique(currentLighted, allLighted)
		for _, pair := range unique {
			lightUp(board, pair)
		}
		allLighted = append(allLighted, unique...)
		currentLighted = findLighted(board)
	}
	return resetLighted(board), len(allLighted)
}

func resetLighted(board [][]int) [][]int {
	for i, line := range board {
		for k, val := range line {
			if val > 9 {
				board[i][k] = 0
			}
		}
	}
	return board
}

func increment(board [][]int) [][]int {
	for i, line := range board {
		for k := range line {
			board[i][k] += 1
		}
	}
	return board
}

func findLighted(board [][]int) []IntPair {
	var result []IntPair
	for i, line := range board {
		for k, val := range line {
			if val > 9 {
				result = append(result, IntPair{k, i})
			}
		}
	}
	return result
}

func lightUp(board [][]int, location IntPair) [][]int {
	if location.x > 0 {
		board[location.y][location.x-1]++
	}
	if location.x < 9 {
		board[location.y][location.x+1]++
	}
	if location.y > 0 {
		board[location.y-1][location.x]++
	}
	if location.y < 9 {
		board[location.y+1][location.x]++
	}
	if location.x > 0 && location.y > 0 {
		board[location.y-1][location.x-1]++
	}
	if location.x < 9 && location.y < 9 {
		board[location.y+1][location.x+1]++
	}
	if location.x > 0 && location.y < 9 {
		board[location.y+1][location.x-1]++
	}
	if location.x < 9 && location.y > 0 {
		board[location.y-1][location.x+1]++
	}
	return board
}

func parseInputDay11(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		var numLine []int
		for _, char := range line {
			numLine = append(numLine, charToInt(char))
		}
		result = append(result, numLine)
	}
	return result
}

func findUnique(listToCheck []IntPair, baseList []IntPair) []IntPair {
	var result []IntPair
	for _, pair := range listToCheck {
		isUnique := true
		for _, base := range baseList {
			if pair.y == base.y && pair.x == base.x {
				isUnique = false
			}
		}
		if isUnique {
			result = append(result, pair)
		}
	}
	return result
}

func printArray(arr [][]int) {
	for _, line := range arr {
		for _, val := range line {
			fmt.Print(val)
		}
		fmt.Println()
	}
	fmt.Println("----------")
}

func isAllFlashed(board [][]int) bool {
	for _, line := range board {
		for _, val := range line {
			if val != 0 {
				return false
			}
		}
	}
	return true
}
