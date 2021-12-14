package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	inputs := readFileByLine("day12.txt")
	var caves []StringPair
	for _, input := range inputs {
		caves = append(caves, parseInputDay12(input))
	}
	paths := buildPaths(caves, []string{}, "start")
	//for _, path := range paths {
	//	for _, cave := range path {
	//		fmt.Print(cave)
	//		fmt.Print(",")
	//	}
	//	fmt.Println()
	//}

	fmt.Println(len(paths))
}

func parseInputDay12(line string) StringPair {
	r := regexp.MustCompile("-")
	split := r.Split(line, -1)
	return StringPair{split[0], split[1]}
}

func buildPaths(caves []StringPair, currentPath []string, currentCave string) [][]string {
	updatedPath := append(currentPath, currentCave)
	if currentCave == "end" {
		return [][]string{updatedPath}
	}
	availableCaves := findAvailableCaves(caves, currentCave)
	var validCaves []string
	for _, cave := range availableCaves {
		if isUpperCase(cave) ||
			!arrayContainsString(updatedPath, cave) ||
			(!arrayContainsLowerCaseDouble(updatedPath) && cave != "start") {
			validCaves = append(validCaves, cave)
		}
	}

	var result [][]string
	for _, cave := range validCaves {
		paths := buildPaths(caves, updatedPath, cave)
		//fmt.Print("current path: [")
		//fmt.Print(updatedPath)
		//fmt.Print("]; next step:")
		//fmt.Print(cave)
		//
		//fmt.Print(": result paths: ")
		//fmt.Println(paths)
		result = append(result, paths...)
		//fmt.Println(result)
	}
	return result
}

func findAvailableCaves(caves []StringPair, currentCave string) []string {
	var result []string
	for _, cave := range caves {
		if cave.first == currentCave && !arrayContainsString(result, cave.second) {
			result = append(result, cave.second)
		} else if cave.second == currentCave && !arrayContainsString(result, cave.first) {
			result = append(result, cave.first)
		}
	}
	return result
}

func arrayContainsLowerCaseDouble(arr []string) bool {
	for i, val1 := range arr {
		if isUpperCase(val1) {
			continue
		}
		for k, val2 := range arr {
			if val1 == val2 && i != k {
				return true
			}
		}
	}
	return false
}

func arrayContainsString(arr []string, toCheck string) bool {
	for _, str := range arr {
		if str == toCheck {
			return true
		}
	}
	return false
}

func isUpperCase(str string) bool {
	return strings.ToUpper(str) == str
}
