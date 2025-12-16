package main

import (
	"fmt"
	"strconv"
	"strings"
)

var day2_sample = []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}

func day2p1(input []string) {
	invalidSum := 0
	var leftChecked = map[string]bool{}
	for r := range strings.SplitSeq(input[0], ",") {
		bounds := strings.Split(r, "-")
		low, _ := strconv.Atoi(bounds[0])
		high, _ := strconv.Atoi(bounds[1])
		for id := low; id <= high; id++ {
			idStr := strconv.Itoa(id)
			size := len(idStr)
			left := idStr[:size/2]
			if leftChecked[left] {
				continue
			}

			doubleNum, _ := strconv.Atoi(strings.Repeat(left, 2))
			if low <= doubleNum && doubleNum <= high {
				leftChecked[left] = true
				invalidSum += doubleNum
			}
		}
	}
	fmt.Printf("day2 part1: %d\n", invalidSum)
}

func day2p2(input []string) {
	invalidSum := 0
	for r := range strings.SplitSeq(input[0], ",") {
		bounds := strings.Split(r, "-")
		low, _ := strconv.Atoi(bounds[0])
		high, _ := strconv.Atoi(bounds[1])
		allIds := map[int]bool{}
		for id := low; id <= high; id++ {
			allIds[id] = true
		}
		for id := low; id <= high; id++ {
			idStr := strconv.Itoa(id)
			size := len(idStr)
			for i := 1; i <= (size/2)+1; i++ {
				left := idStr[:i]
				if size/i == 1 {
					continue
				}
				repeatingStr := strings.Repeat(left, size/i)
				repeating, _ := strconv.Atoi(repeatingStr)
				if allIds[repeating] {
					invalidSum += repeating
					allIds[repeating] = false
					break
				}
			}
		}
	}
	fmt.Printf("day2 part2: %d\n", invalidSum)

}

func day2() {

	input := loadInput("../inputs/day2.txt")
	_ = input
	day2p1(input)
	day2p2(input)
	fmt.Println("finished")

}
