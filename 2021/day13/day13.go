package day13

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(in string, maxX, maxY int) string {
	m := make([][]bool, maxX)
	for x := 0; x < maxX; x++ {
		m[x] = make([]bool, maxY)
	}

	// reading the input
	lines := strings.Split(in, "\n")
	line := lines[0]
	i := 0
	for line != "" {
		l := strings.Split(line, ",")

		m[strToInt(l[0])][strToInt(l[1])] = true

		i++
		line = lines[i]
	}

	// skip empty line
	i++
	l := len("fold along ")

	for ; i < len(lines); i++ {
		f := strings.Split(lines[i][l:], "=")
		foldSize := strToInt(f[1])

		if f[0] == "y" {
			m2 := make([][]bool, maxX)
			// copy existing
			for x := 0; x < maxX; x++ {
				m2[x] = make([]bool, foldSize)

				for y := 0; y < foldSize; y++ {
					m2[x][y] = m[x][y]
				}
			}

			// folding

			for x := range m {
				for y := foldSize; y < maxY; y++ {
					if m[x][y] {
						m2[x][maxY-y-1] = m[x][y]
					}
				}
			}

			maxY = foldSize
			m = m2
		} else {
			m2 := make([][]bool, foldSize)
			// copy existing
			for x := 0; x < foldSize; x++ {
				m2[x] = make([]bool, maxY)

				for y := 0; y < maxY; y++ {
					m2[x][y] = m[x][y]
				}
			}

			// folding

			for x := foldSize; x < maxX; x++ {
				for y := 0; y < maxY; y++ {
					if m[x][y] {
						m2[maxX-x-1][y] = m[x][y]
					}
				}
			}

			maxX = foldSize
			m = m2
		}
	}

	return stringify(m, maxX, maxY)
}

func stringify(m [][]bool, maxX, maxY int) string {
	s := ""
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if m[x][y] {
				s += "█"
			} else {
				s += " "
			}
		}

		s += "\n"
	}

	return s
}

func printMatrix(m [][]bool, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if m[x][y] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
