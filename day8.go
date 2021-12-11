package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := readFileByLine("day8.txt")
	result := 0
	for _, line := range lines {
		firstP, secondP := getInputs(line)
		fmt.Println(firstP)
		mappings := getMappings(firstP)
		num := 0
		for _, str := range secondP {
			num = num*10 + mappings[sorted(str)]
		}
		result += num
	}
	fmt.Println(result)
}

func getInputs(line string) ([]string, []string) {
	delimeterIndex := strings.Index(line, "|")
	firstPart := line[:delimeterIndex-1]
	secondPart := line[delimeterIndex+2:]
	return strings.Split(firstPart, " "), strings.Split(secondPart, " ")
}

func getMappings(examples []string) map[string]int {
	one := findStringsWithLength(examples, 2)[0]
	seven := findStringsWithLength(examples, 3)[0]
	four := findStringsWithLength(examples, 4)[0]
	eight := findStringsWithLength(examples, 7)[0]

	var charsInFour []rune
	for _, char := range four {
		if char != rune(one[0]) && char != rune(one[1]) {
			charsInFour = append(charsInFour, char)
		}
	}
	length5 := findStringsWithLength(examples, 5)
	five := "error"
	three := "error"
	two := "error"
	for _, line := range length5 {
		if containsAll(line, string(charsInFour)) {
			five = line
		} else if containsAll(line, one) {
			three = line
		} else {
			two = line
		}
	}

	length6 := findStringsWithLength(examples, 6)
	nine := "error"
	zero := "error"
	six := "error"
	for _, line := range length6 {
		if containsAll(line, three) {
			nine = line
		} else if containsAll(line, one) {
			zero = line
		} else {
			six = line
		}
	}
	result := make(map[string]int)
	result[sorted(one)] = 1
	result[sorted(two)] = 2
	result[sorted(three)] = 3
	result[sorted(four)] = 4
	result[sorted(five)] = 5
	result[sorted(six)] = 6
	result[sorted(seven)] = 7
	result[sorted(eight)] = 8
	result[sorted(nine)] = 9
	result[sorted(zero)] = 0
	return result
}

func findStringsWithLength(strings []string, length int) []string {
	var result []string
	for _, s := range strings {
		if len(s) == length {
			result = append(result, s)
		}
	}
	return result
}

func containsAll(original string, checker string) bool {
	for _, char := range checker {
		if !strings.ContainsRune(original, char) {
			return false
		}
	}
	return true
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}
