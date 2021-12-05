package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input2.txt
var input2 string

//go:embed input.txt
var input1 string

func main() {
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}
	fmt.Printf("part A: %d\n", partA(input1, matrix))
}

func partA(input string, matrix [][]int) int {
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " -> ")
		l1 := strings.Split(l[0], ",")
		x1 := strToInt(l1[0])
		y1 := strToInt(l1[1])

		l2 := strings.Split(l[1], ",")
		x2 := strToInt(l2[0])
		y2 := strToInt(l2[1])

		calcPaths(matrix, x1, x2, y1, y2)
	}

	return sumMatrix(matrix)
}

func partB(input string, matrix [][]int) int {
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " -> ")
		l1 := strings.Split(l[0], ",")
		x1 := strToInt(l1[0])
		y1 := strToInt(l1[1])

		l2 := strings.Split(l[1], ",")
		x2 := strToInt(l2[0])
		y2 := strToInt(l2[1])

		calcPaths2(matrix, x1, x2, y1, y2)
		// printMatrix(matrix)
	}

	return sumMatrix(matrix)
}

func sumMatrix(matrix [][]int) int {
	sum := 0
	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] > 1 {
				sum++
			}
		}
	}

	return sum
}

func calcPaths(matrix [][]int, x1, x2, y1, y2 int) {
	// horizontal
	if x1 == x2 {
		if y1 > y2 {
			for y := y1; y != (y2 - 1); y += -1 {
				matrix[x1][y]++
			}
		} else {
			for y := y1; y != (y2 + 1); y++ {
				matrix[x1][y]++
			}
		}

		return
	}

	// vertical
	if y1 == y2 {
		if x1 > x2 {
			for x := x1; x != (x2 - 1); x += -1 {
				matrix[x][y1]++
			}
		} else {
			for x := x1; x != (x2 + 1); x++ {
				matrix[x][y1]++
			}
		}

		return
	}
}

func calcPaths2(matrix [][]int, x1, x2, y1, y2 int) {
	// horizontal
	if x1 == x2 {
		if y1 > y2 {
			for y := y1; y != (y2 - 1); y += -1 {
				matrix[x1][y]++
			}
		} else {
			for y := y1; y != (y2 + 1); y++ {
				matrix[x1][y]++
			}
		}

		return
	}

	// vertical
	if y1 == y2 {
		if x1 > x2 {
			for x := x1; x != (x2 - 1); x += -1 {
				matrix[x][y1]++
			}
		} else {
			for x := x1; x != (x2 + 1); x++ {
				matrix[x][y1]++
			}
		}

		return
	}

	// diagonal
	if abs(x1-x2) == abs(y1-y2) {
		diff := abs(x1 - x2)

		if x1 > x2 {
			if y1 > y2 {
				for i := 0; i <= diff; i++ {
					matrix[x1-i][y1-i]++
				}
			} else {
				for i := 0; i <= diff; i++ {
					matrix[x1-i][y1+i]++
				}
			}
		} else {
			if y1 > y2 {
				for i := 0; i <= diff; i++ {
					matrix[x1+i][y1-i]++
				}
			} else {
				for i := 0; i <= diff; i++ {
					matrix[x1+i][y1+i]++
				}
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func printMatrix(m [][]int) {
	for x := range m {
		for y := range m[x] {
			fmt.Printf("%d ", m[x][y])
		}
		fmt.Println()
	}

	fmt.Println()
}
