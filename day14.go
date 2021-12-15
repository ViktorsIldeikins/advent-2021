package main

import (
	"fmt"
)

func main() {
	inputs := readFileByLine("day14.txt")
	polimer := inputs[0]
	transformations := parseAllInputsDay14(inputs)

	for i := 0; i < 40; i++ {
		polimer = doStepPolimer(polimer, transformations)
		fmt.Println(i)
		fmt.Println(len(polimer))
	}

	biggestCount, smallesCount := findSmallestAndBiggestCount(polimer)

	fmt.Println(biggestCount)
	fmt.Println(smallesCount)

	fmt.Println(biggestCount - smallesCount)
}

func findSmallestAndBiggestCount(polimer string) (int, int) {
	biggestCount := 0
	smallesCount := len(polimer) + 1
	for _, char := range polimer {
		count := countChars(polimer, char)
		if count < smallesCount {
			smallesCount = count
		}
		if count > biggestCount {
			biggestCount = count
		}
	}
	return biggestCount, smallesCount
}

func countChars(line string, char int32) int {
	count := 0
	for _, val := range line {
		if val == char {
			count++
		}
	}
	return count
}

func doStepPolimer(polimer string, transformations []StringPair) string {
	result := ""
	for i := 0; i+1 < len(polimer); i++ {
		currentChar := string(polimer[i])
		nextChar := string(polimer[i+1])
		extraChar := findTransformation(transformations, currentChar+nextChar)
		result = result + currentChar + extraChar
	}
	result = result + string(polimer[len(polimer)-1])
	return result
}

func findTransformation(transformation []StringPair, target string) string {
	for _, pair := range transformation {
		if pair.first == target {
			return pair.second
		}
	}
	fmt.Println("ERROR")
	return ""
}

func parseAllInputsDay14(inputs []string) []StringPair {
	var transformations []StringPair
	for i := 2; i < len(inputs); i++ {
		transformations = append(transformations, parseInputDay14(inputs[i]))
	}
	return transformations
}

func parseInputDay14(input string) StringPair {
	return StringPair{input[:2], input[6:7]}
}
