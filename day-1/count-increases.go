package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	inputFile := os.Args[1]

	var windowSize int
	if len(os.Args) < 3 {
		windowSize = 1
	} else {
		input, err := strconv.Atoi(os.Args[2])

		if err != nil {
			log.Fatalf("Invalid input: %s. Please read the README.md!", os.Args[2])
			return
		}

		windowSize = input
	}

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("Not able to open %s. Please read the README.md!", inputFile)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []int

	for scanner.Scan() {
		value, error := strconv.Atoi(scanner.Text())
		if error != nil {
			continue
		} else {
			values = append(values, value)
		}
	}

	lastValue := math.MinInt
	increases := 0
	for last := windowSize - 1; last < len(values); last++ {
		windowValue := 0
		for i := 0; i < windowSize; i++ {
			windowValue = windowValue + values[last-i]
		}
		if lastValue != math.MinInt && windowValue > lastValue {
			increases++
		}
		lastValue = windowValue
	}

	log.Printf("Increases: %d", increases)
}
