package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("No arguments found. Please read the README.md!")
		return
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("Not able to open %s. Please read the README.md!", inputFile)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberOfLines := 0
	var numberOfOnes []int

	for scanner.Scan() {
		line := scanner.Text()

		for pos, char := range line {

			if (len(numberOfOnes) - 1) < pos {
				numberOfOnes = append(numberOfOnes, 0)
			}

			value, err := strconv.Atoi(string(char))

			if err != nil || (value != 0 && value != 1) {
				log.Fatalf("Invalid input line: %s", line)
				return
			}

			if value == 1 {
				numberOfOnes[pos]++
			}
		}

		numberOfLines++
	}

	var gammaRateArray []string
	var epsilonRateArray []string

	for _, ones := range numberOfOnes {
		if ones > numberOfLines/2 {
			gammaRateArray = append(gammaRateArray, "1")
			epsilonRateArray = append(epsilonRateArray, "0")
		} else {
			gammaRateArray = append(gammaRateArray, "0")
			epsilonRateArray = append(epsilonRateArray, "1")
		}
	}

	gammaRate, err := strconv.ParseInt(strings.Join(gammaRateArray[:], ""), 2, 64)
	epsilonRate, err := strconv.ParseInt(strings.Join(epsilonRateArray[:], ""), 2, 64)

	if err != nil {
		log.Fatalf("An error occured when parsing the results: %s, %s", gammaRateArray, epsilonRateArray)
		return
	}

	log.Printf("Gamma Rate: %d", gammaRate)
	log.Printf("Epsilon Rate: %d", epsilonRate)
	log.Printf("Result: %d", gammaRate*epsilonRate)
}
