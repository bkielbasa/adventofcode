package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day2.txt
var input string

func main() {
	partA()
	partB()
}

func partA() {
	horizontal := 0
	depth := 0

	moves := strings.Split(input, "\n")

	for _, move := range moves {
		parts := strings.Split(move, " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}

	fmt.Println("Part A:", depth*horizontal)
}

func partB() {
	horizontal := 0
	depth := 0
	aim := 0

	moves := strings.Split(input, "\n")

	for _, move := range moves {
		parts := strings.Split(move, " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += val
			depth += val * aim
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	fmt.Println("Part B:", depth*horizontal)
}
