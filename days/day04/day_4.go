package day04

import (
	"fmt"

	"github.com/Sp0k/AOC-2025/aoc"
)

func countNeighbors(grid [][]string, targetRow, targetCol int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	count := 0
	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			if row == 0 && col == 0 {
				continue
			}

			nextRow, nextCol := targetRow+row, targetCol+col
			if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
				continue
			}
			if grid[nextRow][nextCol] == "@" {
				count++
			}
		}
	}
	return count
}

type Point struct{ R, C int }

func removeRolls(grid [][]string) (int, [][]string, bool) {
	var rollsToRemove []Point
	sum := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == "@" {
				neighbourCount := countNeighbors(grid, row, col)
				if neighbourCount < 4 {
					sum++
					rollsToRemove = append(rollsToRemove, Point{row, col})
				}
			}
		}
	}

	for _, point := range rollsToRemove {
		grid[point.R][point.C] = "x"
	}

	return sum, grid, len(rollsToRemove) != 0
}

func Solve(input string) {
	grid := aoc.Grid(input)
	
	// Part 1
	sum, grid, removedRolls := removeRolls(grid)
	fmt.Println("[Part 1] There are", sum, "rolls that can be removed by a forklift!")

	// Part 2
	total := sum
	for removedRolls {
		sum, grid, removedRolls = removeRolls(grid)
		total += sum
	}
	fmt.Println("[Part 2] There was", total, "rolls removed!")
}
