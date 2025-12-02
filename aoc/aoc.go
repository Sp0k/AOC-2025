package aoc

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func MustReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func Lines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n")
}

func MustReadLines(path string) []string {
	return Lines(MustReadFile(path))
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatalf("Atoi(%q): %v", s, err)
	}
	return n
}
