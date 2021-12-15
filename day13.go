package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	paperInputs := readFileByLine("day13_1.txt")
	//foldInputs := readFileByLine("day13_2.txt")

	var paper [1500][1500]bool

	for _, input := range paperInputs {
		dot := parsePaperDots(input)
		paper[dot.y][dot.x] = true
	}
	printPaper(paper)
}

func fold(paper [][]bool, index int, foldX bool) [][]bool {
	var result [][]bool
	//shift :=
	if foldX {

	}

	return result
}

func printPaper(paper [1500][1500]bool) {
	for _, line := range paper {
		for _, val := range line {
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parsePaperDots(line string) IntPair {
	r := regexp.MustCompile(",")
	split := r.Split(line, -1)
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return IntPair{x, y}
}
