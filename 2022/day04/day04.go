package day04

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

type section struct {
	begin int
	end   int
}

func (s section) contains(other section) bool {
	return s.begin <= other.begin && s.end >= other.end
}

func (s section) overlaps(other section) bool {
	return s.begin <= other.begin && s.end >= other.begin
}

func partA(s string) int {
	cout := 0

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		sections := strings.Split(line, ",")
		sec1 := extractSection(sections[0])
		sec2 := extractSection(sections[1])

		if sec1.contains(sec2) {
			cout++
			continue
		}

		if sec2.contains(sec1) {
			cout++
			continue
		}
	}

	return cout
}

func extractSection(s string) section {
	parts := strings.Split(s, "-")

	p1, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	p2, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return section{
		begin: p1,
		end:   p2,
	}
}

func partB(s string) int {
	cout := 0

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		sections := strings.Split(line, ",")
		sec1 := extractSection(sections[0])
		sec2 := extractSection(sections[1])

		if sec1.overlaps(sec2) {
			cout++
			continue
		}

		if sec2.overlaps(sec1) {
			cout++
			continue
		}
	}

	return cout
}
