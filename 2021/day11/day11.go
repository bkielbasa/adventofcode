package day11

import (
	_ "embed"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

type octopuses [10][10]byte

func (octs *octopuses) flashing(x, y int) bool {
	return octs[x][y] == ':' // in ascii it's '9' + 1
}

func (octs *octopuses) flash(x, y int) {
	if x > 0 {
		octs[x-1][y]++
		if octs.flashing(x-1, y) {
			octs.flash(x-1, y)
		}
	}

	if y > 0 {
		octs[x][y-1]++
		if octs.flashing(x, y-1) {
			octs.flash(x, y-1)
		}
	}

	if x < 9 {
		octs[x+1][y]++
		if octs.flashing(x+1, y) {
			octs.flash(x+1, y)
		}
	}

	if y < 9 {
		octs[x][y+1]++
		if octs.flashing(x, y+1) {
			octs.flash(x, y+1)
		}
	}

	if x > 0 && y > 0 {
		octs[x-1][y-1]++
		if octs.flashing(x-1, y-1) {
			octs.flash(x-1, y-1)
		}
	}

	if x < 9 && y < 9 {
		octs[x+1][y+1]++
		if octs.flashing(x+1, y+1) {
			octs.flash(x+1, y+1)
		}
	}

	if x < 9 && y > 0 {
		octs[x+1][y-1]++
		if octs.flashing(x+1, y-1) {
			octs.flash(x+1, y-1)
		}
	}

	if x > 0 && y < 9 {
		octs[x-1][y+1]++
		if octs.flashing(x-1, y+1) {
			octs.flash(x-1, y+1)
		}
	}
}

func partA(input string) int {
	octs := octopuses{}

	for x, line := range strings.Split(input, "\n") {
		for y := range line {
			octs[x][y] = line[y]
		}
	}

	flashes := 0

	for i := 0; i < 100; i++ {
		for x := range octs {
			for y := range octs[x] {
				octs[x][y]++

				if octs.flashing(x, y) {
					octs.flash(x, y)
				}
			}
		}

		for x := range octs {
			for y := range octs[x] {
				if octs[x][y] > '9' {
					flashes++
					octs[x][y] = '0'
				}
			}
		}
	}

	return flashes
}

func partB(input string) int {
	octs := octopuses{}

	for x, line := range strings.Split(input, "\n") {
		for y := range line {
			octs[x][y] = line[y]
		}
	}

	for i := 0; ; i++ {
		for x := range octs {
			for y := range octs[x] {
				octs[x][y]++

				if octs.flashing(x, y) {
					octs.flash(x, y)
				}
			}
		}

		flashes := 0
		for x := range octs {
			for y := range octs[x] {
				if octs[x][y] > '9' {
					flashes++
					octs[x][y] = '0'
				}
			}
		}

		if flashes == 100 {
			return i + 1
		}
	}
}
