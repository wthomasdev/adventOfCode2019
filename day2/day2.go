package main

import (
	inputreader "adventOfCode2019/datautils"
	"fmt"
	"math/rand"
)

const stepForward = 4
const breakCode = 99
const expectedOutput = 19690720

func main() {
	numbers := inputreader.Readfile("../data/day02a.txt", ",")
	fmt.Println("answer", calculateNounAndVerb(numbers))
}

func calculateNounAndVerb(numbers []int) int {
	intResult := 0
	noun := 0
	verb := 0
	for intResult != expectedOutput {
		noun = rand.Intn(99)
		verb = rand.Intn(99)
		input := []int{}
		input = append(input, numbers...)
		intResult = intCodeComputer(input, noun, verb)
	}
	return 100*noun + verb
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
