package main

import (
	"fmt"
	"sort"
)

func main() {
	lines := readFileByLine("day10.txt")
	//var completnessChars []int32
	var allScores []int
	for _, line := range lines {
		if !isCorrupted(line) {
			//fmt.Println(line)
			//completnessChars = append(completnessChars, completeLine(line)...)
			//fmt.Println(completnessChars)
			score := 0
			for _, char := range completeLine(line) {
				score = score*5 + mapCharToScore(string(char))
			}
			allScores = append(allScores, score)
		}
	}

	sort.Ints(allScores)
	index := len(allScores) / 2
	fmt.Println(allScores)
	fmt.Println(len(allScores))
	fmt.Println(index)
	fmt.Println(allScores[index])
}

func mapCharToScore(char string) int {
	if char == ")" {
		return 1
	} else if char == "}" {
		return 3
	} else if char == "]" {
		return 2
	} else if char == ">" {
		return 4
	} else {
		return 0
	}
}

func completeLine(line string) []int32 {
	var stack Stack
	for _, element := range line {
		if element == '(' || element == '[' || element == '{' || element == '<' {
			stack.Push(element)
		} else {
			stack.Pop()
		}
	}
	var result []int32
	for !stack.IsEmpty() {
		pop, _ := stack.Pop()
		if pop == '(' {
			result = append(result, ')')
		} else if pop == '[' {
			result = append(result, ']')
		} else if pop == '{' {
			result = append(result, '}')
		} else if pop == '<' {
			result = append(result, '>')
		} else {
			fmt.Println("ERROR")
		}
	}
	return result
}

func isCorrupted(line string) bool {
	return findIllegalChar(line) != ""
}

func findIllegalChar(line string) string {
	var stack Stack
	for _, element := range line {
		if element == '(' || element == '[' || element == '{' || element == '<' {
			stack.Push(element)
		} else {
			if stack.IsEmpty() {
				return string(element)
			}
			pop, _ := stack.Pop()
			if !((pop == '(' && element == ')') ||
				(pop == '[' && element == ']') ||
				(pop == '{' && element == '}') ||
				(pop == '<' && element == '>')) {
				return string(element)
			}
		}
	}
	return ""
}

type Stack []int32

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int32) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
