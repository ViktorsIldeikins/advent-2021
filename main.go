package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text)
		numbers = append(numbers, num)
	}

	count := 0
	for i := 0; i + 3 < len(numbers); i++ {
		if combine(numbers, i) < combine(numbers, i + 1) {
			count++
		}
	}
	fmt.Println(count)
}

func combine(nums []int, firstIndex int) int {
	return nums[firstIndex] + nums[firstIndex + 1] + nums[firstIndex + 2]
}
