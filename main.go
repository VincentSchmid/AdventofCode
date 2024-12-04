package main

import (
	"adventofcode/days"
	"flag"
)

func main() {
	day := flag.Int("day", 1, "the day to run")
	flag.Parse()

	days.RunDay(*day)
}
