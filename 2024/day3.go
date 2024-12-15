package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func loadInput(file string) string {
	f, err := os.Open(file)
	handleErr(err)
	buf, err := io.ReadAll(f)
	handleErr(err)
	return string(buf)
}

func sumSections(input string) int {
	sections := strings.Split(input, "mul(")
	sum := 0
	for _, sec := range sections {
		segments := strings.Split(sec, ",")
		firstInt, err := strconv.Atoi(segments[0])
		if err != nil {
			continue
		}
		last := strings.Split(segments[1], ")")[0]
		if len(last) > 3 {
			continue
		}
		lastInt, err := strconv.Atoi(last)
		handleErr(err)
		sum += firstInt * lastInt
	}
	return sum
}

func day3part1() {
	input := strings.ReplaceAll(loadInput("day3.txt"), "\n", "")
	print(sumSections(input))
}

func day3part2() {
	input := strings.ReplaceAll(loadInput("day3.txt"), "\n", "")
	dos := strings.Split(input, "do()")
	sum := 0
	for _, d := range dos {
		donts := strings.Split(d, "don't()")
		sum += sumSections(donts[0])
	}
	fmt.Println("sum:", sum)
}
