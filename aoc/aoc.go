package aoc

import (
	"log"
	"os"
	"strconv"
	"strings"
)

//
// I/O
//

// Reads the entire file
func MustReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// Splits a string into lines
func Lines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimSpace(s)
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n")
}

func Columns(s string) [][]string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimSpace(s)

	rawLines := Lines(s)

	var rows [][]string
	for _, line := range rawLines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		rows = append(rows, fields)
	}

	if len(rows) == 0 {
		return [][]string{}
	}

	colCount := len(rows[0])
	cols := make([][]string, colCount)

	for _, row := range rows {
		for i, v := range row {
			cols[i] = append(cols[i], v)
		}
	}

	return cols
}

func Grid(s string) [][]string {
	lines := Lines(s)
	if lines == nil {
		return nil
	}

	grid := make([][]string, len(lines))
	for i, line := range lines {
		row := make([]string, len(line))
		for j := 0; j < len(line); j++ {
			row[j] = string(line[j])
		}
		grid[i] = row
	}
	return grid
}

// Reads the entire file and splits it into lines
func MustReadLines(path string) []string {
	return Lines(MustReadFile(path))
}

// Splits string into blocks separated by 1 or more lines
func SplitBlocks(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n\n")
}

// Splits string into values from CSV string
func CSVStrings(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimSpace(s)
	if (s == "") {
		return nil
	}
	return strings.Split(s, ",")
}

//
// Parsing
//

// Parses an int
func MustAtoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatalf("Atoi(%q): %v", s, err)
	}
	return n
}

// Parse an array of strings into ints
func ParseArr(arr []string) []int {
	parsed := make([]int, len(arr))
	for i, n := range arr {
		parsed[i] = MustAtoi(n)
	}
	return parsed
}

// Parses whitespace-separated ints from a line
func FieldsInt(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i] = MustAtoi(p)
	}
	return nums
}

// Parses comma-separated ints from a line
func CSVInts(line string) []int {
	parts := strings.Split(strings.TrimSpace(line), ",")
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		nums = append(nums, MustAtoi(p))
	}
	return nums
}

// Parses an int into a string
func IntToStr(n int) string {
	if n == 0 {
		return "0"
	}

	negative := n < 0

	var buf [20]byte
	i := len(buf)

	for nn := n; nn != 0; nn /= 10 {
		digit := nn % 10
		if digit < 0 {
			digit = -digit
		}
		i--
		buf[i] = byte('0' + digit)
	}

	if negative {
		i--
		buf[i] = '-'
	}

	return string(buf[i:])
}

//
// Math
//

// Returns the sum of a slice of ints
func SumInts(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

//
// Go funcs
//

// Removes a value from a slice
func Remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Inserts a value in ascending order inside the slice
func InsertInOrder(slice []int, value int) []int {
	i := 0
	for i < len(slice) && slice[i] <= value {
		i++
	}

	slice = append(slice, 0)
	copy(slice[i+1:], slice[i:])
	slice[i] = value
	return slice
}

// Checks if a slice contains a specific value
func Contains(slice []int, value int) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}
