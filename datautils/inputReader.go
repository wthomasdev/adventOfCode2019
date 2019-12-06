// Package inputreader includes utils for reading data in from files
package inputreader

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Readfile takes in a path and returns an int array
func Readfile(filepath string, separator string) []int {
	file, err := ioutil.ReadFile(filepath)
	sanitizedFile := strings.TrimSpace(string(file))
	if err != nil {
		log.Fatalf("could not read file %e", err)
	}
	nStrings := strings.Split(sanitizedFile, separator)
	numbers := []int{}
	for _, nS := range nStrings {
		i, err := strconv.Atoi(nS)
		if err != nil {
			log.Fatalf("Error converting string to number: %e", err)
		}
		numbers = append(numbers, i)
	}
	return numbers
}
