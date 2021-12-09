package day9

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

type board struct {
	m       [][]int
	sizeX   int
	sizeY   int
	visited [][]bool
}

func partA(input string, sizex, sizey int) int {
	matrix := make([][]int, sizex)
	for i := range matrix {
		matrix[i] = make([]int, sizey)
	}

	risk := 0
	lines := strings.Split(input, "\n")

	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			matrix[x][y] = int(lines[x][y]) - 48
		}
	}

	for x := range matrix {
		for y := range matrix[x] {
			if x < sizex-1 {
				if matrix[x+1][y] <= matrix[x][y] {
					continue
				}
			}

			if y < sizey-1 {
				if matrix[x][y+1] <= matrix[x][y] {
					continue
				}
			}

			if x > 0 {
				if matrix[x-1][y] <= matrix[x][y] {
					continue
				}
			}

			if y > 0 {
				if matrix[x][y-1] <= matrix[x][y] {
					continue
				}
			}

			risk += matrix[x][y] + 1
		}
	}

	return risk
}

func partB(input string, sizex, sizey int) int {
	visited := make([][]bool, sizex)
	matrix := make([][]int, sizex)
	for i := range matrix {
		matrix[i] = make([]int, sizey)
		visited[i] = make([]bool, sizey)
	}

	lines := strings.Split(input, "\n")

	for x := 0; x < sizex; x++ {
		for y := 0; y < sizey; y++ {
			matrix[x][y] = int(lines[x][y]) - 48
		}
	}

	b := board{m: matrix, visited: visited, sizeX: sizex, sizeY: sizey}

	basinsSize := []int{}
	for x := range matrix {
		for y := range matrix[x] {
			if x < sizex-1 {
				if matrix[x+1][y] <= matrix[x][y] {
					continue
				}
			}

			if y < sizey-1 {
				if matrix[x][y+1] <= matrix[x][y] {
					continue
				}
			}

			if x > 0 {
				if matrix[x-1][y] <= matrix[x][y] {
					continue
				}
			}

			if y > 0 {
				if matrix[x][y-1] <= matrix[x][y] {
					continue
				}
			}

			size := b.calcBasinSize(x, y)
			basinsSize = append(basinsSize, size)
		}
	}

	sort.Ints(basinsSize)
	s := len(basinsSize)

	return basinsSize[s-1] * basinsSize[s-2] * basinsSize[s-3]
}

func (b *board) calcBasinSize(x, y int) int {
	if b.m[x][y] == 9 {
		return 0
	}

	if b.visited[x][y] {
		return 0
	}

	b.visited[x][y] = true

	size := 1

	if x < b.sizeX-1 {
		size += b.calcBasinSize(x+1, y)
	}

	if y < b.sizeY-1 {
		size += b.calcBasinSize(x, y+1)
	}

	if x > 0 {
		size += b.calcBasinSize(x-1, y)
	}

	if y > 0 {
		size += b.calcBasinSize(x, y-1)
	}

	return size
}
