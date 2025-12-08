package day08

import (
	"fmt"
	"math"
	"sort"

	"github.com/Sp0k/AOC-2025/aoc"
)

type JunctionBox struct {
	X, Y, Z int
}

type DSU struct {
	parent []int
	size   []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(a, b int) {
	ra := d.find(a)
	rb := d.find(b)
	if ra == rb {
		return
	}

	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
}

func distance(box1, box2 JunctionBox) float64 {
	dx := float64(box1.X - box2.X)
	dy := float64(box1.Y - box2.Y)
	dz := float64(box1.Z - box2.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type edge struct {
	i, j int
	dist float64
}

func productOfThreeLargestComponents(dsu *DSU, n int) int {
	componentSize := make(map[int]int)
	for i := range n {
		root := dsu.find(i)
		componentSize[root]++
	}

	var sizes []int
	for _, s := range componentSize {
		sizes = append(sizes, s)
	}

	if len(sizes) < 3 {
		return 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func Solve(input string) {
	lines := aoc.Lines(input)

	var boxes []JunctionBox
	for _, line := range lines {
		if line == "" {
			continue
		}
		pos := aoc.CSVInts(line)
		boxes = append(boxes, JunctionBox{
			X: pos[0],
			Y: pos[1],
			Z: pos[2],
		})
	}

	n := len(boxes)
	if n == 0 {
		fmt.Println("no boxes")
		return
	}

	var edges []edge
	for i := range n {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{
				i:    i,
				j:    j,
				dist: distance(boxes[i], boxes[j]),
			})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].dist < edges[b].dist
	})

	dsu := newDSU(n)

	const connectionsForPart1 = 1000 
	components := n
	connectionsConsidered := 0
	part1Done := false
	var part1Result int

	var lastEdge edge
	lastEdgeSet := false

	for _, e := range edges {
		if dsu.find(e.i) != dsu.find(e.j) {
			dsu.union(e.i, e.j)
			components--

			if components == 1 && !lastEdgeSet {
				lastEdge = e
				lastEdgeSet = true
			}
		}

		connectionsConsidered++

		if !part1Done && connectionsConsidered == connectionsForPart1 {
			part1Result = productOfThreeLargestComponents(dsu, n)
			part1Done = true
		}

		if components == 1 {
			break
		}
	}

	if part1Done {
		fmt.Println("[Part 1] The product of the three biggest circuits is", part1Result)
	} else {
		fmt.Println("[Part 1] Not enough connections considered")
	}

	if !lastEdgeSet {
		fmt.Println("[Part 2] Could not find a final connecting edge")
		return
	}

	x1 := boxes[lastEdge.i].X
	x2 := boxes[lastEdge.j].X
	product := x1 * x2
	fmt.Println("[Part 2] The product of X coordinates of the last connection is", product)
}
