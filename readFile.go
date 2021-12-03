package main

import (
	"bufio"
	"os"
)

func readFileByLine(fileName string) []string {
	file, _ := os.Open("inputs/" + fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
