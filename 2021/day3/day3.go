package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day3.txt
var input string

func main() {
	partA()
	partB()
}

func partA() {
	ones := [12]int{}
	lines := strings.Split(input, "\n")
	inputSize := len(lines)

	for _, line := range lines {
		for i, ch := range line {
			if ch == '1' {
				ones[i]++
			}
		}
	}

	gammaBytes := make([]byte, 12)
	epsilonBytes := make([]byte, 12)
	for i := range ones {
		if ones[i] > (inputSize / 2) {
			gammaBytes[i] = '1'
			epsilonBytes[i] = '0'
		} else {
			gammaBytes[i] = '0'
			epsilonBytes[i] = '1'
		}
	}

	gamma, _ := strconv.ParseInt(string(gammaBytes), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonBytes), 2, 64)

	fmt.Printf("gamma: %d, epsilon: %d, solution: %d\n", gamma, epsilon, gamma*epsilon)
}

func partB() {
	oxygen := strings.Split(input, "\n")
	co2 := make([]string, len(oxygen))
	copy(co2, oxygen)

	i := 0
	for {
		common := mostCommon(oxygen, i)
		oxygen = onlyContainingByte(oxygen, common, i)

		if len(oxygen) == 1 {
			break
		}

		i++
	}

	i = 0
	for {
		common := leastCommon(co2, i)
		co2 = onlyContainingByte(co2, common, i)

		if len(co2) == 1 {
			break
		}

		i++
	}

	ox, _ := strconv.ParseInt(oxygen[0], 2, 64)
	c, _ := strconv.ParseInt(co2[0], 2, 64)

	fmt.Println("part B: ", ox, c, "=", ox*c)
}

func leastCommon(input []string, position int) byte {
	c := map[byte]int{
		'0': 0,
		'1': 0,
	}

	for i := range input {
		c[input[i][position]]++
	}

	// if there's more 'one's
	if c['1'] > c['0'] {
		return '0'
	}

	return '1'
}

func mostCommon(input []string, position int) byte {
	c := map[byte]int{
		'0': 0,
		'1': 0,
	}

	for i := range input {
		c[input[i][position]]++
	}

	// if there's more 'one's
	if c['1'] >= c['0'] {
		return '1'
	}

	return '0'
}

func onlyContainingByte(input []string, val byte, position int) []string {
	filtered := []string{}

	for _, line := range input {
		if line[position] == val {
			filtered = append(filtered, line)
		}
	}

	return filtered
}
