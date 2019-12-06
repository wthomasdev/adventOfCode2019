package main

import (
	inputreader "adventOfCode2019/datautils"
	"fmt"
)

const stepForward = 4

func main() {
	numbers := inputreader.Readfile("../data/day02a.txt", ",")
	numbers[1] = 12
	numbers[2] = 2

	fmt.Println("answer", calculateOpCode(numbers, 1, 2, 12, 2))
}

func calculateOpCode(numbers []int, modifierloc1 int, modifierloc2 int, modifierNum1 int, modifierNum2 int) int {
	numbers[modifierloc1] = modifierNum1
	numbers[modifierloc2] = modifierNum2
	for i := 0; i < len(numbers); i += stepForward {
		if numbers[i] == 99 {
			break
		}
		opCode, pos1, pos2, resultPointer := numbers[i], numbers[i+1], numbers[i+2], numbers[i+3]
		switch opCode {
		case 1:
			numbers[resultPointer] = numbers[pos1] + numbers[pos2]
		case 2:
			numbers[resultPointer] = numbers[pos1] * numbers[pos2]
		}
	}
	return numbers[0]
}
