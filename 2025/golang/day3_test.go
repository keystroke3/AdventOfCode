package main

import "testing"

var input = loadInput("../inputs/day3.txt")

func BenchmarkDay3p1v1(b *testing.B) {
	for b.Loop() {
		day3p1v1(input)
	}
}

func BenchmarkDay3p1v2(b *testing.B) {
	for b.Loop() {
		day3p1v2(input)
	}
}

func BenchmarkDay3p1v3(b *testing.B) {
	for b.Loop() {
		day3p1v3(input)
	}
}
