package main

import (
	"strconv"
)

var day3sample = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func day3p1v1(input []string) int {
	sumOfJoltages := 0
	for _, bank := range input {
		highestJoltage := 0
		excludeLast := bank[:len(bank)-1]
		for i, leftJolt := range excludeLast {
			excludeFirst := bank[i+1:]
			for _, rightJolt := range excludeFirst {
				joltage, _ := strconv.Atoi(string([]rune{leftJolt, rightJolt}))
				if joltage > highestJoltage {
					highestJoltage = joltage
				}
			}
			if leftJolt == '9' {
				break
			}
		}
		sumOfJoltages += highestJoltage
	}
	return sumOfJoltages
}

func day3p1v2(input []string) int {
	sumOfJoltages := 0
	for _, bank := range input {
		highestJoltage := 0
		highestLeftIndex := 0
		highestLeft := '0'

		excludeLast := bank[:len(bank)-1]
		for i, leftJolt := range excludeLast {
			if leftJolt > highestLeft {
				highestLeft = leftJolt
				highestLeftIndex = i
			}
			if leftJolt == '9' {
				break
			}
		}
		excludeFirstI := bank[highestLeftIndex+1:]
		for _, rightJolt := range excludeFirstI {
			joltage, _ := strconv.Atoi(string([]rune{highestLeft, rightJolt}))
			if joltage >= highestJoltage {
				highestJoltage = joltage
			}
		}
		sumOfJoltages += highestJoltage
	}
	return sumOfJoltages
}

func day3p1v3(input []string) int {
	sumOfJoltages := 0
	for _, bank := range input {
		highestLeftIndex := 0
		highestLeft := '0'

		excludeLast := bank[:len(bank)-1]
		for i, leftJolt := range excludeLast {
			if leftJolt > highestLeft {
				highestLeft = leftJolt
				highestLeftIndex = i
			}
			if leftJolt == '9' {
				break
			}
		}
		excludeFirstI := bank[highestLeftIndex+1:]
		highestRight := '0'
		for _, rightJolt := range excludeFirstI {
			if rightJolt >= highestRight {
				highestRight = rightJolt
			}
			if rightJolt == '9' {
				break
			}
		}
		joltage, _ := strconv.Atoi(string([]rune{highestLeft, highestRight}))
		sumOfJoltages += joltage
	}
	return sumOfJoltages
}

func day3p2(input []string) int {
	sumOfJoltages := 0
	for _, bank := range input {
		batteries := 12
		start := 0
		heighestJolts := make([]rune, batteries)
		for i := batteries - 1; i >= 0; i-- {
			trimmed := bank[start : len(bank)-i]
			hiJolt := '0'
			hiJoltIndex := 0
			for index, jolt := range trimmed {
				if jolt > hiJolt {
					hiJolt = jolt
					hiJoltIndex = index
				}
			}
			start = start + 1 + hiJoltIndex
			heighestJolts[(batteries-1)-i] = hiJolt
		}
		joltageStr := string(heighestJolts)
		joltage, _ := strconv.Atoi(joltageStr)
		sumOfJoltages += joltage
	}
	return sumOfJoltages
}

func day3() {
	input := loadInput("../inputs/day3.txt")
	println("day3 part1:", day3p1v3(input), "\nday3 part 2:", day3p2(input))
}
