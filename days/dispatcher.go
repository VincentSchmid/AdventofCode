package days

import (
	"adventofcode/days/day01"
	"adventofcode/days/day02"
	"adventofcode/days/day03"
	"log"
)

func RunDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	default:
		log.Fatalf("Day %d is not implemented.", day)
	}
}
