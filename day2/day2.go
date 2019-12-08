package main

import (
	inputreader "adventOfCode2019/datautils"
	"fmt"
)

const stepForward = 4
const breakCode = 99
const expectedOutput = 19690720

func main() {
	numbers := inputreader.ReadNumberfile("../data/day02a.txt", ",")
	answer, _ := calculateNounAndVerb(numbers)
	fmt.Printf("Answer: %v \n", answer)
}

func calculateNounAndVerb(numbers []int) (int, error) {
	for n := 0; n < 99; n++ {
		for v := 0; v < 99; v++ {
			input := append([]int{}, numbers...)
			outut := intCodeComputer(input, n, v)
			if outut == expectedOutput {
				return 100*n + v, nil
			}
		}
	}
	return 0, fmt.Errorf("Could not find the answer")
}

func intCodeComputer(numbers []int, noun, verb int) int {
	numbers[1] = noun
	numbers[2] = verb
	for i := 0; i < len(numbers); i += stepForward {
		if numbers[i] == breakCode {
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
