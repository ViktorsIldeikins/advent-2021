package main

import (
	"regexp"
	"strconv"
)

func parseInput(input string) []int {
	r := regexp.MustCompile(",")
	split := r.Split(input, -1)
	var result []int
	for _, num := range split {
		parsedNum, _ := strconv.Atoi(num)
		result = append(result, parsedNum)
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}