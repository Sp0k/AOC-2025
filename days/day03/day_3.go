package day03

import (
	"fmt"

	"github.com/Sp0k/AOC-2025/aoc"
)

func parseLine(line string) []int {
	var parsedLine []int

	for _, c := range line {
		parsedLine = append(parsedLine, aoc.MustAtoi(string(c)))
	}

	return parsedLine
}

func findBestKIndices(digits []int, k int) []int {
	n := len(digits)
	if k >= n {
		res := make([]int, n)
		for i := range res {
			res[i] = i
		}
		return res
	}

	toRemove := n - k
	stack := make([]int, 0, n)

    for i, d := range digits {
        for toRemove > 0 && len(stack) > 0 && digits[stack[len(stack)-1]] < d {
            stack = stack[:len(stack)-1]
            toRemove--
        }
        stack = append(stack, i)
    }

	if len(stack) > k {
		stack = stack[:k]
	}

	return stack
}

func Solve(input string) {
	banks := aoc.Lines(input)

	sum1 := 0
	sum2 := 0
	for _, bank := range banks {
		parsedBank := parseLine(bank)
		
		idx2 := findBestKIndices(parsedBank, 2)
		var num2Str string
		for _, idx := range idx2 {
			num2Str += aoc.IntToStr(parsedBank[idx])
		}
		sum1 += aoc.MustAtoi(num2Str)

		idx12 := findBestKIndices(parsedBank, 12)
		var fullJoltage string
		for _, idx := range idx12 {
			fullJoltage += aoc.IntToStr(parsedBank[idx])
		}
		sum2 += aoc.MustAtoi(fullJoltage)
	}
	fmt.Println("[Part 1] The total output joltage is", sum1)
	fmt.Println("[Part 2] The total output joltage is", sum2)
}
