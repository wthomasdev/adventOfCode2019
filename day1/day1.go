package main

import (
	"adventOfCode2019/datautils"
	"fmt"
)

func main() {
	numbers := inputreader.Readfile("../data/day01a.txt")
	fmt.Println(calculateTotalFuel(numbers))
}

func calculateModuleFuel(mass int, sum int) int {
	fuelReq := (mass / 3) - 2
	if fuelReq <= 0 {
		return sum
	}
	newSum := sum + fuelReq
	return calculateModuleFuel(fuelReq, newSum)
}

func calculateTotalFuel(massArr []int) int {
	sum := 0
	for _, mass := range massArr {
		sum += calculateModuleFuel(mass, 0)
	}
	return sum
}
