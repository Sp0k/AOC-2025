package day01

import (
	"fmt"

	"github.com/Sp0k/AOC-2025/aoc"
)

var dialVal int
var zeroCounterP1 int
var zeroCounterP2 int

func turnLeft(rotations int) {
	if dialVal == 0 {
		zeroCounterP2 += rotations / 100
	} else if rotations >= dialVal {
		zeroCounterP2 += 1 + (rotations - dialVal)/100
	}
	dialVal = ((dialVal - rotations)%100 + 100) % 100
}

func turnRight(rotations int) {
	zeroCounterP2 += (dialVal + rotations) / 100
	dialVal = (dialVal + rotations) % 100
}

func handleRotationInstruction(instruction string) {
	rotation := aoc.MustAtoi(instruction[1:])

	if rune(instruction[0]) == 'L' {
		turnLeft(rotation)
	} else {
		turnRight(rotation)
	}
}

func Solve(input string) {
	instructions := aoc.Lines(input)
	dialVal = 50
	zeroCounterP1 = 0
	zeroCounterP2 = 0

	for i := range(instructions) {
		handleRotationInstruction(instructions[i])
		if (dialVal == 0) {
			zeroCounterP1++
		}
	}

	fmt.Println("First code: ", zeroCounterP1)
	fmt.Println("Second code: ", zeroCounterP2)
}
