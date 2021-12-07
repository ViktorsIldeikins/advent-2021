package main

import "fmt"

func main() {
	positions := parseInput(readFileByLine("day7.txt")[0])

	smallestFuel := 1000000000
	biggestPos := 0
	for i := range positions {
		if biggestPos < positions[i] {
			biggestPos = positions[i]
		}
	}

	for i := 0; i < biggestPos; i++ {
		fuel := calculateTotalFuel(positions, i)
		if smallestFuel > fuel {
			smallestFuel = fuel
		}
	}

	fmt.Println(smallestFuel)
}

func calculateTotalFuel(positions []int, target int) int {
	totalFuel := 0
	for i := range positions {
		distance := abs(positions[i] - target)
		totalFuel += calculateFuelFromDistance(distance)
	}
	return totalFuel
}

func calculateFuelFromDistance(distance int) int {
	result := 0
	for i := 1; i <= distance; i++ {
		result += i
	}
	return result
}
