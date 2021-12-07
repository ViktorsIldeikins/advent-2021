package main

import (
	"fmt"
)

const SPLIT_AGE = 7

func main() {
	//subjects := []int {1}
	//subjects := parseInput(readFileByLine("day6.txt")[0])

	//for i:= 0; i< 256; i ++ {
	//	subjects = ageOneDay(subjects)
	//}
	//fmt.Println(producedPopulation(1, 256) * 84)
	//fmt.Println(producedPopulation(2, 256) * 59)
	//fmt.Println(producedPopulation(3, 256) * 54)
	//fmt.Println(producedPopulation(4, 256) * 48)
	//fmt.Println(producedPopulation(5, 256) * 55)

	fmt.Println(521372966772 + 331408259732 + 281730055068 + 226852841952 + 240252760495)
}

func producedPopulation(daysUntilSplit int, daysForCalculation int) int {
	if daysForCalculation == 0 || daysForCalculation < daysUntilSplit {
		return 1
	}

	if daysUntilSplit > 0 {
		return producedPopulation(0, daysForCalculation-daysUntilSplit)
	}

	return producedPopulation(6, daysForCalculation-1) + producedPopulation(8, daysForCalculation-1)
}

func ageOneDay(subjects []int) []int {
	newSubjects := 0
	for i := range subjects {
		subjects[i]--
		if subjects[i] == -1 {
			subjects[i] += SPLIT_AGE
			newSubjects++
		}
	}
	for i := 0; i < newSubjects; i++ {
		subjects = append(subjects, 8)
	}
	return subjects
}
