package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadSaftyReport() [][]int {
	f, err := os.Open("day2.txt")
	handleErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	var reports [][]int
	for _, line := range lines {
		splitl := strings.Split(line, " ")
		// fmt.Printf("%q\n", splitl)
		report := make([]int, len(splitl))
		for i, l := range splitl {
			n, err := strconv.Atoi(l)
			handleErr(err)
			report[i] = n
		}
		reports = append(reports, report)
	}
	return reports
}

func day2part1() {
	reports := loadSaftyReport()
	safeReports := 0
	for _, report := range reports {
		increasing := report[1] > report[0]
		safe := true
		for i := 0; i < len(report)-1; i++ {
			a := report[i]
			b := report[i+1]
			if a == b {
				safe = false
				break
			}
			if increasing && a > b {
				safe = false
				break
			}
			if !increasing && a < b {
				safe = false
				break
			}
			diff := absoluteDiff(a, b)
			if diff > 3 {
				safe = false
				break
			}
		}
		if safe {
			safeReports += 1
		}
	}
	fmt.Println(safeReports)
}

func isSafe(report []int) bool {
	increasing := report[1] > report[0]
	safe := true
	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]
		if a == b {
			return false
		}
		if increasing && a > b {
			return false
		}
		if !increasing && a < b {
			return false
		}
		diff := absoluteDiff(a, b)
		if diff > 3 {
			return false
		}
	}
	return safe
}

func removeElement(slice []int, index int) []int {
	var report []int
	for i, n := range slice {
		if i == index {
			continue
		}
		report = append(report, n)
	}
	return report
}

func day2part2() {
	reports := loadSaftyReport()
	safeReports := 0
	for _, report := range reports {
		safe := isSafe(report)
		if safe {
			safeReports += 1
			continue
		}
		for i := range report {
			_report := removeElement(report, i)
			safe = isSafe(_report)
			if safe {
				safeReports += 1
				break
			}
		}

	}
	fmt.Println(safeReports)
}
