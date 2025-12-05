package day05

import (
	"fmt"
	"strings"

	"github.com/Sp0k/AOC-2025/aoc"
)

type Range struct { Min, Max int}

func sortRanges(ranges []Range) []Range {
	for i := 1; i < len(ranges); i++ {
		key := ranges[i]
		j := i - 1
		for j >= 0 && ranges[j].Min > key.Min {
			ranges[j+1] = ranges[j]
			j--
		}
		ranges[j+1] = key
	}

	return ranges
}

func calculateFreshIngredients(ranges []Range, ids []int) int {
	count := 0
	for _, id := range ids {
		fresh := false
		for _, r := range ranges {
			if id >= r.Min && id <= r.Max {
				fresh = true
				break
			}
		}
		if fresh {
			count++
		}
	}
	return count
}

func calculateAvailableFreshIds(ranges []Range) int {
	ranges = sortRanges(ranges)

	var merged []Range
	currMin := ranges[0].Min
	currMax := ranges[0].Max
	for i := 1; i < len(ranges); i++ {
		if ranges[i].Min <= currMax + 1 {
			if ranges[i].Max > currMax {
				currMax = ranges[i].Max
			}
		} else {
			merged = append(merged, Range{Min: currMin, Max: currMax})
			currMin = ranges[i].Min
			currMax = ranges[i].Max
		}
	}	
	merged = append(merged, Range{Min: currMin, Max: currMax})

	total := 0
	for _, r := range merged {
		total += r.Max - r.Min + 1
	}

	return total
}

func Solve(input string) {
	lists := aoc.SplitBlocks(input)
	strRanges := aoc.Lines(lists[0])
	strIds := aoc.Lines(lists[1])

	ids := aoc.ParseArr(strIds)
	ranges := make([]Range, len(strRanges))
	for i, r := range strRanges {
		vals := strings.Split(r, "-")
		min := aoc.MustAtoi(vals[0])
		max := aoc.MustAtoi(vals[1])
		ranges[i] = Range{Min: min, Max: max}
	}

	part1 := calculateFreshIngredients(ranges, ids)
	part2 := calculateAvailableFreshIds(ranges)

	fmt.Println("[Part 1] There is", part1, "spoiled ingredients!")
	fmt.Println("[Part 2] There is", part2, "fresh ingredient id available!")
}
