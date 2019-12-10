package main

import (
	"fmt"
	"strconv"
	"strings"
)

const rangeStart = 130254
const rangeFinish = 678275

func main() {
	validPasswords := calculateValidPasswords(rangeStart, rangeFinish)
	fmt.Println("ValidPasswords", validPasswords)
}

func calculateValidPasswords(start, finish int) int {
	validPasswords := 0
	for i := start; i <= finish; i++ {
		stringVal := strconv.Itoa(i)
		checkedPassword := validatePassword(stringVal)
		if checkedPassword {
			validPasswords++
		}
	}
	return validPasswords
}

func validatePassword(password string) bool {
	passwordArr := strings.Split(password, "")
	valid := checkAdjacentDigitAndIncrease(passwordArr)
	return valid
}

func checkAdjacentDigitAndIncrease(arr []string) bool {
	adjacent := false
	increasing := false
	for idx, char := range arr {
		if idx != 0 {
			charInt, _ := strconv.Atoi(char)
			prevCharInt, _ := strconv.Atoi(arr[idx-1])
			if charInt >= prevCharInt {
				increasing = true
			} else {
				increasing = false
				break
			}
		}
		if len(arr)-1 == idx {
			break
		}
		if char == arr[idx+1] {
			adjacent = true
		}
	}
	return adjacent && increasing
}
