package day06

import (
	"fmt"
	"strings"

	"github.com/Sp0k/AOC-2025/aoc"
)

type EquationType int

const (
	Multiplication EquationType = iota
	Addition
)

func findEquationType(symbol byte) EquationType {
	if symbol == '*' {
		return Multiplication
	}
	return Addition
}

func eval(nums []int, op EquationType) int {
	if len(nums) == 0 {
		return 0
	}

	if op == Multiplication {
		prod := 1
		for _, n := range nums {
			prod *= n
		}
		return prod
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func part1(worksheet string) int {
	columns := aoc.Columns(worksheet)

	total := 0
	for _, col := range columns {
		if len(col) == 0 {
			continue
		}

		op := findEquationType(col[len(col)-1][0])

		nums := make([]int, len(col)-1)
		for i := 0; i < len(col)-1; i++ {
			nums[i] = aoc.MustAtoi(col[i])
		}

		total += eval(nums, op)
	}

	return total
}

func buildGrid(worksheet string) ([]string, int) {
	worksheet = strings.ReplaceAll(worksheet, "\r\n", "\n")
	worksheet = strings.TrimRight(worksheet, "\n")

	lines := strings.Split(worksheet, "\n")

	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	for i, line := range lines {
		if len(line) < width {
			lines[i] = line + strings.Repeat(" ", width-len(line))
		}
	}

	return lines, width
}

func part2(worksheet string) int {
	worksheet = strings.ReplaceAll(worksheet, "\t", " ")

	lines, width := buildGrid(worksheet)
	height := len(lines)

	total := 0
	opChar := byte('+')
	var nums []int
	currentDigits := make([]byte, 0)

	flushProblem := func() {
		if len(nums) == 0 {
			return
		}
		total += eval(nums, findEquationType(opChar))
		nums = nums[:0]
	}

	for x := range width {
		if len(currentDigits) > 0 {
			n := aoc.MustAtoi(string(currentDigits))
			nums = append(nums, n)
			currentDigits = currentDigits[:0]
		}

		colBlank := true

		for y := range height {
			ch := lines[y][x]
			if ch == ' ' {
				continue
			}
			colBlank = false

			if ch == '+' || ch == '*' {
				opChar = ch
				continue
			}

			currentDigits = append(currentDigits, ch)
		}

		if colBlank {
			if len(currentDigits) > 0 {
				n := aoc.MustAtoi(string(currentDigits))
				nums = append(nums, n)
				currentDigits = currentDigits[:0]
			}
			flushProblem()
		}
	}

	if len(currentDigits) > 0 {
		n := aoc.MustAtoi(string(currentDigits))
		nums = append(nums, n)
	}
	flushProblem()

	return total
}

func Solve(input string) {
	sum1 := part1(input)
	sum2 := part2(input)

	fmt.Println("[Part 1] The total is", sum1)
	fmt.Println("[Part 2] The total is", sum2)
}
