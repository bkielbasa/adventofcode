package day01

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(s string) int {
	elfs := []int{}
	sum := 0
	for _, line := range strings.Split(s, "\n") {

		if line == "" {
			elfs = append(elfs, sum)
			sum = 0
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += i
	}

	if sum > 0 {
		elfs = append(elfs, sum)
	}

	max := elfs[0]

	for _, e := range elfs {
		if e > max {
			max = e
		}
	}

	return max
}

func partB(s string) int {
	elfs := []int{}
	sum := 0
	for _, line := range strings.Split(s, "\n") {

		if line == "" {
			elfs = append(elfs, sum)
			sum = 0
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		sum += i
	}

	if sum > 0 {
		elfs = append(elfs, sum)
	}

	sort.Ints(elfs)
	l := len(elfs)

	return elfs[l-1] + elfs[l-2] + elfs[l-3]
}
