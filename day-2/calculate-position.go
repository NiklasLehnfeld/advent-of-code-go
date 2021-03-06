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

	aim := 0
	depth := 0
	horizontal := 0

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")

		errorHappened := false

		if len(command) != 2 ||
			(command[0] != "forward" &&
				command[0] != "up" &&
				command[0] != "down") {
			errorHappened = true
		}

		direction := command[0]
		number, err := strconv.Atoi(command[1])

		if err != nil {
			errorHappened = true
		}

		if errorHappened {
			log.Fatalf("Input contains invalid command: %s. Please read README.md!", line)
			return
		}

		switch direction {
		case "forward":
			horizontal += number
			depth += aim * number
		case "up":
			aim -= number
		case "down":
			aim += number
		}
	}

	log.Printf("Aim: %d", aim)
	log.Printf("Depth: %d", depth)
	log.Printf("Horizontal: %d\n", horizontal)

	log.Printf("Solution Part 1: %d", aim*horizontal)
	log.Printf("Solution Part 2: %d", depth*horizontal)
}
