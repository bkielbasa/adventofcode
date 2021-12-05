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
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " -> ")
		l1 := strings.Split(l[0], ",")
		x1 := strToInt(l1[0])
		y1 := strToInt(l1[1])

		l2 := strings.Split(l[1], ",")
		x2 := strToInt(l2[0])
		y2 := strToInt(l2[1])

		sum += calcPaths(matrix, x1, x2, y1, y2)
	}

	return sum
}

func partB(input string, matrix [][]int) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " -> ")
		l1 := strings.Split(l[0], ",")
		x1 := strToInt(l1[0])
		y1 := strToInt(l1[1])

		l2 := strings.Split(l[1], ",")
		x2 := strToInt(l2[0])
		y2 := strToInt(l2[1])

		sum += calcPaths2(matrix, x1, x2, y1, y2)
	}

	return sum
}

func calcPaths(matrix [][]int, x1, x2, y1, y2 int) int {
	c := 0
	// horizontal
	if x1 == x2 {
		if y1 > y2 {
			for y := y1; y != (y2 - 1); y += -1 {
				matrix[x1][y]++
				if matrix[x1][y] == 2 {
					c++
				}
			}
		} else {
			for y := y1; y != (y2 + 1); y++ {
				matrix[x1][y]++
				if matrix[x1][y] == 2 {
					c++
				}
			}
		}

		return c
	}

	// vertical
	if y1 == y2 {
		if x1 > x2 {
			for x := x1; x != (x2 - 1); x += -1 {
				matrix[x][y1]++
				if matrix[x][y1] == 2 {
					c++
				}
			}
		} else {
			for x := x1; x != (x2 + 1); x++ {
				matrix[x][y1]++
				if matrix[x][y1] == 2 {
					c++
				}
			}
		}
	}

	return c
}

func calcPaths2(matrix [][]int, x1, x2, y1, y2 int) int {
	c := 0
	// horizontal
	if x1 == x2 {
		if y1 > y2 {
			for y := y1; y != (y2 - 1); y += -1 {
				matrix[x1][y]++
				if matrix[x1][y] == 2 {
					c++
				}
			}
		} else {
			for y := y1; y != (y2 + 1); y++ {
				matrix[x1][y]++
				if matrix[x1][y] == 2 {
					c++
				}
			}
		}

		return c
	}

	// vertical
	if y1 == y2 {
		if x1 > x2 {
			for x := x1; x != (x2 - 1); x += -1 {
				matrix[x][y1]++
				if matrix[x][y1] == 2 {
					c++
				}
			}
		} else {
			for x := x1; x != (x2 + 1); x++ {
				matrix[x][y1]++
				if matrix[x][y1] == 2 {
					c++
				}
			}
		}

		return c
	}

	// diagonal
	if abs(x1-x2) == abs(y1-y2) {
		diff := abs(x1 - x2)

		if x1 > x2 {
			if y1 > y2 {
				for i := 0; i <= diff; i++ {
					matrix[x1-i][y1-i]++
					if matrix[x1-i][y1-i] == 2 {
						c++
					}
				}
			} else {
				for i := 0; i <= diff; i++ {
					matrix[x1-i][y1+i]++
					if matrix[x1-i][y1+i] == 2 {
						c++
					}
				}
			}
		} else {
			if y1 > y2 {
				for i := 0; i <= diff; i++ {
					matrix[x1+i][y1-i]++
					if matrix[x1+i][y1-i] == 2 {
						c++
					}
				}
			} else {
				for i := 0; i <= diff; i++ {
					matrix[x1+i][y1+i]++
					if matrix[x1+i][y1+i] == 2 {
						c++
					}
				}
			}
		}
	}

	return c
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
