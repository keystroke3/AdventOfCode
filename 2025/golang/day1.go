package main

import (
	"fmt"
	"log"
	"strconv"
)

const maxPositions = 100

func part1(input []string) {
	currentPos := 50
	zeros := 0
	for _, rotation := range input {
		clicks, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal(err)
		}
		if rotation[0] == 'L' {
			currentPos = maxPositions - ((currentPos-clicks)*-1)%maxPositions
		} else {
			currentPos = (currentPos + clicks) % maxPositions
		}
		if currentPos == 0 || currentPos == maxPositions {
			zeros++
		}
	}
	fmt.Printf("day 1 part 1: %d\n", zeros)
}

func part2(input []string) {
	currentPos := 50
	zeros := 0
	for _, rotation := range input {
		clicks, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal(err)
		}
		if rotation[0] == 'L' {
			for range clicks {
				currentPos--
				if currentPos == 0 {
					zeros++
				}
				if currentPos < 0 {
					currentPos = maxPositions - (currentPos * -1)
				}
			}
		} else {
			for range clicks {
				currentPos++
				if currentPos == 0 {
					zeros++
				}
				if currentPos >= maxPositions {
					currentPos = maxPositions - currentPos
					zeros++
				}
			}
		}
	}
	fmt.Printf("day 1 part 2: %d\n", zeros)
}

func day1() {
	input := loadInput("../inputs/day1.txt")
	part1(input)
	part2(input)
}
