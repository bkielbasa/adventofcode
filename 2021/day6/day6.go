package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input2.txt
var inputSmall string

//go:embed input.txt
var input1 string

func partA(input string, iterations int) int {
	fishes := []int{}

	for _, f := range strings.Split(input, ",") {
		fishes = append(fishes, strToInt(f))
	}

	for i := 0; i < iterations; i++ {
		toAdd := []int{}
		for j := 0; j < len(fishes); j++ {
			if fishes[j] == 0 {
				toAdd = append(toAdd, 8)
				fishes[j] = 6
			} else {
				fishes[j]--
			}
		}

		fishes = append(fishes, toAdd...)
	}

	return len(fishes)
}

func partB(input string, iterations int) int {
	fishes := []int{}

	for _, f := range strings.Split(input, ",") {
		fishes = append(fishes, strToInt(f))
	}

	result := make(map[int]int)

	for _, f := range fishes {
		result[f] += 1
	}

	for d := 0; d < iterations; d++ {
		created := make(map[int]int)

		for i := 0; i <= 8; i++ {
			k := i
			v, ok := result[k]
			if !ok {
				continue
			}
			if k == 0 {
				delete(result, k)
				created[6] += v
				created[8] += v
				continue
			}
			result[k] -= v
			result[k-1] += v
		}

		for k, v := range created {
			result[k] += v
		}
	}

	var total int
	for _, v := range result {
		total += v
	}
	return total
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
