package main

var solutions = map[int]func(){
	1: day1,
	2: day2,
	3: day3,
}

func main() {
	day := 3
	solutions[day]()
}
