package day02

import (
	_ "embed"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(s string) int {
	points := map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
	}

	mappings := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	sum := 0

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		p1 := line[0:1]
		p2 := line[2:3]

		p2 = mappings[p2]
		p1, p2 = p2, p1

		if p1 == p2 {
			sum += points[p1] + 3
			continue
		}

		if p1 == "A" && p2 == "B" {
			sum += points[p1]
			continue
		}

		if p1 == "A" && p2 == "C" {
			sum += points[p1] + 6
			continue
		}

		if p1 == "B" && p2 == "A" {
			sum += points[p1] + 6
			continue
		}

		if p1 == "B" && p2 == "C" {
			sum += points[p1]
			continue
		}

		if p1 == "C" && p2 == "A" {
			sum += points[p1]
			continue
		}

		if p1 == "C" && p2 == "B" {
			sum += points[p1] + 6
			continue
		}
	}

	return sum
}

func partB(s string) int {
	points := map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
	}

	sum := 0

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		p1 := line[0:1]
		p2 := line[2:3]

		if p1 == "A" { // rock
			if p2 == "X" { // loses
				sum += points["C"]
				continue
			}

			if p2 == "Y" { // draft
				sum += points[p1] + 3
				continue
			}

			if p2 == "Z" { // wins
				sum += points["B"] + 6
				continue
			}
		}

		if p1 == "B" { // paper
			if p2 == "X" { // loses
				sum += points["A"]
				continue
			}

			if p2 == "Y" { // draft
				sum += points[p1] + 3
				continue
			}

			if p2 == "Z" { // wins
				sum += points["C"] + 6
				continue
			}
		}

		if p1 == "C" { // scissors
			if p2 == "X" { // loses
				sum += points["B"]
				continue
			}

			if p2 == "Y" { // draft
				sum += points[p1] + 3
				continue
			}

			if p2 == "Z" { // wins
				sum += points["A"] + 6
				continue
			}
		}
	}

	return sum
}
