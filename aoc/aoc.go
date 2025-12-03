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
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n")
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
