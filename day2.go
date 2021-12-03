package main

import (
	"fmt"
	"strings"
)

func main() {
	inputs := readFileByLine("day2.txt")
	fmt.Println(len(inputs))

	horizontal := 0
	vertical := 0
	aim := 0
	for _, input := range inputs {
		amount := input[len(input)-1]
		if strings.Contains(input, "forward") {
			horizontal += int(amount - '0')
			vertical += int(amount-'0') * aim
		} else if strings.Contains(input, "down") {
			aim += int(amount - '0')
		} else if strings.Contains(input, "up") {
			aim -= int(amount - '0')
		} else {
			fmt.Println("couldn't parse:" + input)
		}
	}

	fmt.Println(horizontal * vertical)
}
