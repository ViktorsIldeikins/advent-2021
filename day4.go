package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputLines := readFileByLine("day4.txt")
	drawnNumbers := getNumbersFromString(inputLines[0], ",")
	var boards [][][]int
	for i := 1; i < len(inputLines); i += 6 {
		fmt.Println(i)
		boards = append(boards, parseBoard(inputLines[i:i+6]))
	}

	fmt.Println("----")
	biggestDraws := -1
	var boardNum int
	var numOfDraws int
	for index, board := range boards {
		draws := numOfDrawsForBoard(drawnNumbers, board)
		fmt.Println(draws)
		if draws > biggestDraws {
			biggestDraws = draws
			boardNum = index
			numOfDraws = draws
		}
	}

	fmt.Println(boardNum + 1)

	sumOfBoard := sumOfBoard(drawnNumbers, boards[boardNum], numOfDraws)
	lastNum := drawnNumbers[numOfDraws]
	fmt.Println(sumOfBoard)
	fmt.Println(lastNum)
	fmt.Println(sumOfBoard * lastNum)

	//fmt.Println(drawnNumbers)
	//for _, board := range boards {
	//	for _, ints := range board {
	//		fmt.Println(ints)
	//	}
	//	fmt.Println("")
	//}
}

func sumOfBoard(drawnNumbers []int, board [][]int, numberOfDraws int) int {
	sum := 0
	for idx, row := range board {
		for idx2 := range row {
			if findIndex(drawnNumbers, board[idx][idx2]) > numberOfDraws {
				sum += board[idx][idx2]
			}
		}
	}
	return sum
}

func numOfDrawsForBoard(drawnNumbers []int, board [][]int) int {
	smallest := len(drawnNumbers) + 1
	for _, row := range board {
		draws := numOfDrawsForRow(drawnNumbers, row)
		if draws < smallest {
			smallest = draws
		}
	}

	for i := 0; i < len(board); i++ {
		var column []int
		for k := 0; k < len(board); k++ {
			column = append(column, board[k][i])
		}
		draws := numOfDrawsForRow(drawnNumbers, column)
		if draws < smallest {
			smallest = draws
		}
	}

	return smallest
}

func numOfDrawsForRow(drawnNumbers []int, numbersToFind []int) int {
	biggest := 0
	for _, num := range numbersToFind {
		index := findIndex(drawnNumbers, num)
		if biggest < index {
			biggest = index
		}
	}
	return biggest
}

func parseBoard(inputLines []string) [][]int {
	var result [][]int
	for i := 1; i <= 5; i++ {
		result = append(result, getNumbersFromString(inputLines[i], " {1,2}"))
	}
	return result
}

func getNumbersFromString(inputLine string, regex string) []int {
	r := regexp.MustCompile(regex)
	split := r.Split(strings.TrimSpace(inputLine), -1)
	var result []int
	for _, num := range split {
		parsedNum, _ := strconv.Atoi(num)
		result = append(result, parsedNum)
	}
	return result
}

func findIndex(arr []int, toFind int) int {
	for index, value := range arr {
		if value == toFind {
			return index
		}
	}
	return -1
}
