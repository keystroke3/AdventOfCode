package main

import (
	"bufio"
	"log"
	"os"
)

func loadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) > 0 {
			lines = append(lines, scanner.Text())
		}

	}
	return lines
}
