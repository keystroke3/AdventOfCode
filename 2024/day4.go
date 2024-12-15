package main

import (
	"fmt"
	"strings"
)

func countHorizontal(input string) int {
	foward := strings.Count(input, "XMAS")
	backward := strings.Count(input, "SAMX")
	return foward + backward
}

func countVertical(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for i := 0; i < len(lines)-4; i++ {
		for j := range lines[i] {
			str := fmt.Sprintf("%s%s%s%s", string(lines[i][j]), string(lines[i+1][j]), string(lines[i+2][j]), string(lines[i+3][j]))
			if str == "XMAS" || str == "SAMX" {
				total += 1
			}
		}
	}
	return total
}

func countDiagonal(input string) int {
	lines := strings.Split(input, "\n")
	totalForward := 0
	totalBackward := 0
	for i := 0; i < len(lines)-4; i++ {
		for j := 0; j < len(lines[i])-3; j++ {
			str := fmt.Sprintf("%s%s%s%s", string(lines[i][j]), string(lines[i+1][j+1]), string(lines[i+2][j+2]), string(lines[i+3][j+3]))
			if str == "XMAS" || str == "SAMX" {
				totalForward += 1
			}
		}
	}

	for i := 0; i < len(lines)-4; i++ {
		for j := 3; j < len(lines[i]); j++ {
			str := fmt.Sprintf("%s%s%s%s", string(lines[i][j]), string(lines[i+1][j-1]), string(lines[i+2][j-2]), string(lines[i+3][j-3]))
			if str == "XMAS" || str == "SAMX" {
				totalBackward += 1
			}
		}
	}
	return totalBackward + totalForward
}

func day4part1() {
	input := loadInput("day4.txt")
	hor := countHorizontal(input)
	vert := countVertical(input)
	diag := countDiagonal(input)
	total := hor + vert + diag
	fmt.Println("total:", total)
}

func day4part2() {
	input := loadInput("day4.txt")
	lines := strings.Split(input, "\n")
	total := 0
	for i := 0; i < len(lines)-3; i++ {
		for j := 0; j < len(lines[0])-2; j++ {
			forward := fmt.Sprintf("%s%s%s", string(lines[i][j]), string(lines[i+1][j+1]), string(lines[i+2][j+2]))
			backward := fmt.Sprintf("%s%s%s", string(lines[i][j+2]), string(lines[i+1][j+1]), string(lines[i+2][j]))
			if (forward == "MAS" || forward == "SAM") && (backward == "MAS" || backward == "SAM") {
				total += 1
			}

		}
	}

	fmt.Println("total:", total)
}
