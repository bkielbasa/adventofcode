package day03

import (
	_ "embed"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(s string) int {
	sum := 0

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		p1 := line[0 : len(line)/2]
		p2 := line[len(line)/2:]

		for _, c := range p1 {
			if strings.Contains(p2, string(c)) {
				if c >= 'a' {
					sum += int(c) - 96
				} else {
					sum += int(c) - 64 + 26
				}
				break
			}
		}
	}

	return sum
}

func contains(s string, c string) bool {
	for _, v := range s {
		if string(v) == c {
			return true
		}
	}
	return false
}

func partB(s string) int {
	sum := 0

	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines); i += 3 {
		if lines[i] == "" {
			continue
		}
		p1 := lines[i]
		p2 := lines[i+1]
		p3 := lines[i+2]

		m1 := map[rune]int{}
		m2 := map[rune]int{}
		m3 := map[rune]int{}

		for _, c := range p1 {
			m1[c]++
		}

		for _, c := range p2 {
			m2[c]++
		}

		for _, c := range p3 {
			m3[c]++
		}

		for _, c := range p1 {
			if m2[c] > 0 && m3[c] > 0 {
				if c >= 'a' {
					sum += int(c) - 96
				} else {
					sum += int(c) - 64 + 26
				}
				break
			}
		}
	}

	return sum
}
