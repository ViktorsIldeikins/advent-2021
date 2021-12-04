package main

import "fmt"

func main() {
	inputs := readFileByLine("day3.txt")
	//fmt.Println(inputs)
	part2(inputs)
}

func part2(inputs []string) {
	correct1 := getCorrect(inputs, 0, true)
	correct2 := getCorrect(inputs, 0, false)

	fmt.Println(toDecimal(correct1))
	fmt.Println(toDecimal(correct2))
	fmt.Println(toDecimal(correct2) * toDecimal(correct1))
}

func toDecimal(input string) int {
	result := 0
	for _, bit := range input {
		result = result*2 + int(bit-'0')
	}
	return result
}

func getCorrect(inputs []string, indexToCheck int, modeMostCommon bool) string {
	if len(inputs) == 1 {
		return inputs[0]
	} else if len(inputs) == 0 {
		fmt.Println("Problems!!")
		return "0"
	}
	size := len(inputs)
	counts := countOnes(inputs)
	amountOfOnes := counts[indexToCheck]
	var bitToKeep uint8
	if (modeMostCommon && amountOfOnes >= size-amountOfOnes) ||
		(!modeMostCommon && amountOfOnes < size-amountOfOnes) {
		bitToKeep = '1'
	} else {
		bitToKeep = '0'
	}
	var filteredStrings []string
	for _, input := range inputs {
		if input[indexToCheck] == bitToKeep {
			filteredStrings = append(filteredStrings, input)
		}
	}
	return getCorrect(filteredStrings, indexToCheck+1, modeMostCommon)
}

func part1(inputs []string) {
	counter := countOnes(inputs)

	size := len(inputs)
	fmt.Println(counter)
	first := 0
	second := 0
	for _, bit := range counter {
		first = first << 1
		second = second << 1
		if bit > (size - bit) {
			first++
		} else {
			second++
		}
	}
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(first * second)
}

func countOnes(inputs []string) [12]int {
	counter := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, input := range inputs {
		for i := 0; i < len(input); i++ {
			if input[i] == '1' {
				counter[i]++
			}
		}
	}
	return counter
}
