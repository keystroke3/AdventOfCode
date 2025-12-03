package main

var solutions = map[int]func(){
	1: day1,
}

func main() {
	day := 1
	solutions[day]()
}
