package days

import (
	"log"
    "adventofcode/days/day01"
    "adventofcode/days/day02"
    // Import other days as needed
)

func RunDay(day int) {
    switch day {
    case 1:
        day01.Run()
    case 2:
        day02.Run()
    default:
        log.Fatalf("Day %d is not implemented.", day)
    }
}
