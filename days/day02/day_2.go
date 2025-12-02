package day02

import (
	"fmt"
	"strings"

	"github.com/Sp0k/AOC-2025/aoc"
)

func part1(a, b int) int {
	sum := 0

	for n := a; n <= b; n++ {
		str := aoc.IntToStr(n)
		if len(str) % 2 != 0 {
			continue
		}

		halfSize := len(str) / 2

		if str[0:halfSize] == str[halfSize:] {
			sum += n
		}
	}

	return sum
}

func IsInvalidID(n int) bool {
    str := aoc.IntToStr(n)
    numLen := len(str)

    for patternLen := 1; patternLen <= numLen/2; patternLen++ {
        if numLen%patternLen != 0 {
            continue
        }

        pattern := str[0:patternLen]
        times := numLen / patternLen

        if times < 2 {
            continue
        }

        repeatedOk := true

        for i := 1; i < times; i++ {
            start := i * patternLen
            end := start + patternLen
            block := str[start:end]

            if block != pattern {
                repeatedOk = false
                break
            }
        }

        if repeatedOk {
            return true
        }
    }

    return false
}

func part2(a, b int) int {
	sum := 0

	for n := a; n <= b; n++ {
		if (IsInvalidID(n)) {
			fmt.Println(n)
			sum += n
		}
	}

	return sum
}

func Solve(input string) {
	ranges := aoc.CSVStrings(input)

	sum1 := 0
	sum2 := 0

	for _, r := range ranges {
		rangeSplit := strings.Split(r, "-")
		a := aoc.MustAtoi(rangeSplit[0])
		b := aoc.MustAtoi(rangeSplit[1])

		sum1 += part1(a, b)
		sum2 += part2(a, b)
	}

	fmt.Println("[Part 1] The invalid ids produce", sum1)
	fmt.Println("[Part 2] The invalid ids produce", sum2)
}
