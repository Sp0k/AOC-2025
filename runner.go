package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Sp0k/AOC-2025/aoc"
	d01 "github.com/Sp0k/AOC-2025/days/day01"
	d02 "github.com/Sp0k/AOC-2025/days/day02"
	d03 "github.com/Sp0k/AOC-2025/days/day03"
	d04 "github.com/Sp0k/AOC-2025/days/day04"
	d05 "github.com/Sp0k/AOC-2025/days/day05"
	d06 "github.com/Sp0k/AOC-2025/days/day06"
	// TODO: Import next days here
)

type Solver func(input string)

var solvers = map[int]Solver{
	1: d01.Solve,
	2: d02.Solve,
	3: d03.Solve,
	4: d04.Solve,
	5: d05.Solve,
	6: d06.Solve,
	// TODO: Add next solvers here
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <day> <input-file>", os.Args[0])
	}

	dayStr := os.Args[1]
	inputPath := os.Args[2]

	day, err := strconv.Atoi(dayStr)
	if err != nil || day <= 0 {
		log.Fatalf("invalid day %q (must be a positive integer)", dayStr)
	}

	solver, ok := solvers[day]
	if !ok {
		log.Fatalf("no solver registered for day %d", day)
	}

	input := aoc.MustReadFile(inputPath)

	solver(input)
}
