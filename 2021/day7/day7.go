package day7

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input2.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(input string) int {
	craps := stringsToInts(strings.Split(input, ","))
	min := 9999999999

	ch := make(chan int, len(craps)-2)

	for i := 2; i <= len(craps); i++ {
		go func(i int) {
			ch <- alignTo(craps, i)
		}(i)
	}

	for i := 2; i <= len(craps); i++ {
		c := <-ch

		if c < min {
			min = c
		}
	}

	return min
}

var results map[int]int

func partB(input string) int {
	craps := stringsToInts(strings.Split(input, ","))
	min := 9999999999

	for i := 2; i <= len(craps); i++ {
		c := alignToB(craps, i)
		if c < min {
			min = c
		}
	}

	return min
}

func alignTo(craps []int, position int) int {
	c := 0

	for _, crap := range craps {
		c += abs(position - crap)
	}

	return c
}

func alignToB(craps []int, position int) int {
	c := 0

	for _, crap := range craps {
		moves := abs(position - crap)

		val, has := results[moves]
		if has {
			c += val
			continue
		}

		sum := 0
		for i := 1; i <= moves; i++ {
			sum += i
		}

		results[moves] = sum
		c += sum
	}

	return c
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
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
