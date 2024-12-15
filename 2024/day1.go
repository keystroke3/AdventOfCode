package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func absoluteDiff(i int, j int) int {
	dif := i - j
	if dif < 0 {
		return -dif
	}
	return dif
}

func sortInput(s [][]string) ([]int, []int) {
	left := make([]int, len(s))
	right := make([]int, len(s))
	for i, pair := range s {
		l, _ := strconv.Atoi(pair[0])
		r, _ := strconv.Atoi(pair[1])
		left[i] = l
		right[i] = r
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func loadCsvInput() [][]string {
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	input, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return input

}

func day1part1() {
	input := loadCsvInput()
	left, right := sortInput(input)
	sum := 0
	for i, n := range left {
		sum += int(math.Abs(float64(n - right[i])))
	}
	fmt.Println("total difference:", sum)
}

func day1part2() {
	input := loadCsvInput()
	left, right := sortInput(input)
	rightFrequency := make(map[int]int)
	for _, n := range right {
		count, ok := rightFrequency[n]
		if ok {
			count += 1
			rightFrequency[n] = count
			continue
		}
		rightFrequency[n] = 1
	}
	score := 0
	for _, n := range left {
		if count, ok := rightFrequency[n]; ok {
			score += count * n
		}
	}
	fmt.Println(score)

}
