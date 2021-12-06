package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	board := [1000][1000]int{}
	inputs := readFileByLine("day5.txt")

	counter := 0

	for _, line := range inputs {
		x1, y1, x2, y2 := getCoordinates(line)
		//fmt.Println(x1, y1, x2, y2)
		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				board[x1][i]++
				if board[x1][i] == 2 {
					counter++
				}
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				board[i][y1]++
				if board[i][y1] == 2 {
					counter++
				}
			}
		} else {
			var xUpdate int
			var yUpdate int
			if x1 < x2 {
				xUpdate = 1
			} else {
				xUpdate = -1
			}
			if y1 < y2 {
				yUpdate = 1
			} else {
				yUpdate = -1
			}

			k := y1
			for i := x1; i != x2+xUpdate; i += xUpdate {
				board[i][k]++
				if board[i][k] == 2 {
					counter++
				}
				k += yUpdate
			}
		}
		//printBoard(board)

	}

	fmt.Println(counter)

}

func printBoard(board [10][10]int) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func getCoordinates(input string) (x1 int, y1 int, x2 int, y2 int) {
	r := regexp.MustCompile("\\d+")
	//r.FindStringSubmatch()
	split := r.FindAllString(strings.TrimSpace(input), -1)
	x1, _ = strconv.Atoi(split[0])
	y1, _ = strconv.Atoi(split[1])
	x2, _ = strconv.Atoi(split[2])
	y2, _ = strconv.Atoi(split[3])
	return x1, y1, x2, y2
}
