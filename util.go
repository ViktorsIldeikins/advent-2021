package main

import (
	"regexp"
	"strconv"
)

func parseInputUtil(input string) []int {
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

func charToInt(char int32) int {
	return int(char - '0')
}

type StringPair struct {
	first  string
	second string
}
