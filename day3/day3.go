package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Point is the x,y location a wire has travelled
type Point struct {
	x int
	y int
}

// WireGrid stores the locations both wires have travelled
type WireGrid map[Point]bool

// NewWireGrid initializes the map with 0,0 as starting point
func NewWireGrid() WireGrid {
	centralPort := Point{0, 0}
	wg := make(WireGrid)
	wg[centralPort] = true
	return wg
}

func (wg WireGrid) calculateAndAddPoints(instructions string, lastPosition Point) {
	instructionsSplit := strings.Split(instructions, ",")
	firstInstruction := string(instructionsSplit[0])
	direction := string(firstInstruction[0])
	distance, _ := strconv.Atoi(string(firstInstruction[1:]))
	if direction == "R" {
		for i := 1; i <= distance; i++ {
			newDistance := lastPosition.y + i
			newPoint := Point{lastPosition.x, newDistance}
			wg[newPoint] = true
			fmt.Println("new point", newPoint)
			fmt.Println(wg)
		}
	}
}

func main() {
	instructions := ReadAndSeparateWireInstructions()
	firstWire := instructions[0]
	fmt.Println("firstWire", firstWire)
	// secondWire := instructions[1]
	wireGrid := NewWireGrid()
	firstInstruction := firstWire[0]
	fmt.Println("firstInstruction", firstInstruction)
	wireGrid.calculateAndAddPoints(firstWire[0], Point{0, 0})
}

// ReadAndSeparateWireInstructions specific function for this challenge, splits into two lists of instructions
func ReadAndSeparateWireInstructions() [][]string {
	file, err := os.Open("../data/day03a.txt")
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines [][]string

	for scanner.Scan() {
		lines = append(lines, []string{scanner.Text()})
	}
	return lines
}
