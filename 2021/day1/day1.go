package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day1.txt
var input string

func main() {
	partA()
	partB()
}

func partB() {
	counter := 0
	measurements := stringsToInts(strings.Split(input, "\n"))

	for i := 3; i < len(measurements); i++ {
		if (measurements[i-1] + measurements[i-2] + measurements[i-3]) < (measurements[i-0] + measurements[i-1] + measurements[i-2]) {
			counter++
		}
	}

	fmt.Println("Part B:", counter)
}

func partA() {
	counter := 0
	measurements := stringsToInts(strings.Split(input, "\n"))

	for i := 1; i < len(measurements); i++ {
		if measurements[i-1] < measurements[i] {
			counter++
		}
	}

	fmt.Println("Part A: ", counter)
}

func stringsToInts(s []string) []int {
	ints := []int{}
	for i := 0; i < len(s); i++ {
		mes, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, mes)
	}
	return ints
}
