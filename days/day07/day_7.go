package day07

import (
	"fmt"

	"github.com/Sp0k/AOC-2025/aoc"
)

type Point struct { X, Y int }

type Beam struct {
	CurrPos Point
	Active bool
	Journeys int
}

func printGraph(graph [][]string) {
	for _, line := range graph {
		for _, ch := range line {
			fmt.Print(ch)
		}
		fmt.Println()
	}
	fmt.Println()
}

func allBeamInactive(beams []Beam) bool {
	for _, beam := range beams {
		if beam.Active == true {
			return false
		}
	}
	return true
}

func part1(grid [][]string, startPos Point) int {
	var beams []Beam
	startingBeam := Beam{Active: true, CurrPos: startPos}
	beams = append(beams, startingBeam)
	totalSplits := 0
	for !allBeamInactive(beams) {
		for i := range beams {
			if !beams[i].Active {
				continue
			}

			if beams[i].CurrPos.Y + 1 >= len(grid) || grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X] == "x" {
				beams[i].Active = false 
			} else if grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X] == "^" {
				beams[i].Active = false
				totalSplits++

				if grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X - 1] != "|" && grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X - 1] != "^" {
					leftBeam := Beam{CurrPos: Point{X: beams[i].CurrPos.X - 1, Y: beams[i].CurrPos.Y + 1}, Active: true}
					beams = append(beams, leftBeam)
					grid[leftBeam.CurrPos.Y][leftBeam.CurrPos.X] = "|"
				}
				if grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X + 1] != "|" && grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X + 1] != "^" {
					rightBeam := Beam{CurrPos: Point{X: beams[i].CurrPos.X + 1, Y: beams[i].CurrPos.Y + 1}, Active: true}
					beams = append(beams, rightBeam)
					grid[rightBeam.CurrPos.Y][rightBeam.CurrPos.X] = "|"
				}

				grid[beams[i].CurrPos.Y + 1][beams[i].CurrPos.X] = "x"
			} else {
				beams[i].CurrPos.Y++
				grid[beams[i].CurrPos.Y][beams[i].CurrPos.X] = "|"
			}
		}
	}
	printGraph(grid)
	return totalSplits
}

func followTrachyon(grid [][]string, currPos Point, memo [][]int) int {
	if currPos.X < 0 || currPos.X >= len(grid[0]) {
		return 1
	}

	if memo[currPos.Y][currPos.X] != -1 {
		return memo[currPos.Y][currPos.X]
	}

	y := currPos.Y
	x := currPos.X

	for y+1 < len(grid) && grid[y+1][x] != "^" {
		y++
	}

	var res int
	if y+1 >= len(grid) {
		res = 1
	} else {
		res = followTrachyon(grid, Point{X: x - 1, Y: y + 1}, memo) +
		followTrachyon(grid, Point{X: x + 1, Y: y + 1}, memo)
	}

	memo[currPos.Y][currPos.X] = res
	return res
}

func part2(grid [][]string, startPos Point) int {
	memo := make([][]int, len(grid))
	for i := range memo {
		memo[i] = make([]int, len(grid[0]))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return followTrachyon(grid, startPos, memo)
}

func Solve(input string) {
	grid := aoc.Grid(input)
	var startPos Point

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "S" {
				startPos = Point{X: j, Y: i + 1}
				grid[startPos.Y][startPos.X] = "|"
			}
		}
	}

	part2 := part2(grid, startPos)
	part1 := part1(grid, startPos)

	fmt.Println("[Part 1] There was", part1, "total splits!")
	fmt.Println("[Part 2] There is", part2, "potential timelines!")
}
