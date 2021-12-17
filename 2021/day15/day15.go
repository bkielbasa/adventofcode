package day15

import (
	"container/heap"
	_ "embed"
	"math"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(in string) int {
	lines := strings.Split(in, "\n")
	m := len(lines)
	n := len(lines[0])
	var grid [100][100]byte
	var seen [100][100]bool
	var dist [100][100]uint16
	for i, row := range lines {
		for j := range row {
			grid[i][j] = lines[i][j] - '0'
			dist[i][j] = math.MaxUint16
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// Keep a min heap of distances
	h := make(minHeap, 1, 100)
	h[0] = pos{0, 0, 0}

	// While there are entries in the min heap (always true for this)
	for {
		x := heap.Pop(&h).(pos)
		seen[x.i][x.j] = true
		if x.i == m-1 && x.j == n-1 {
			return int(x.val)
		}
		for _, nei := range [][2]int{
			{x.i + 1, x.j}, {x.i - 1, x.j}, {x.i, x.j - 1}, {x.i, x.j + 1},
		} {
			ii, jj := nei[0], nei[1]
			if !ok(ii, jj) || seen[ii][jj] {
				continue
			}
			risk := x.val + uint16(grid[ii][jj])
			if risk >= dist[ii][jj] {
				continue
			}
			dist[ii][jj] = risk
			heap.Push(&h, pos{ii, jj, risk})
		}
	}
}

func partB(in string) int {
	lines := strings.Split(in, "\n")
	m := len(lines)
	n := len(lines[0])
	var grid [500][500]byte
	var seen [500][500]bool
	var dist [500][500]uint16
	for i, row := range lines {
		for j := range row {
			// Copy board + adjust values
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					ii, jj := k*m+i, l*n+j
					grid[ii][jj] = lines[i][j] - '0' + byte(k+l)
					dist[ii][jj] = math.MaxUint16
					if grid[ii][jj] > 9 {
						grid[ii][jj] %= 9
					}
				}
			}
		}
	}
	m = 5 * m
	n = 5 * n
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// Keep a min heap of distances
	h := make(minHeap, 1, 100)
	h[0] = pos{0, 0, 0}

	// While there are entries in the min heap (always true for this)
	for {
		x := heap.Pop(&h).(pos)
		seen[x.i][x.j] = true
		if x.i == m-1 && x.j == n-1 {
			return int(x.val)
		}
		for _, nei := range [][2]int{
			{x.i + 1, x.j}, {x.i - 1, x.j}, {x.i, x.j - 1}, {x.i, x.j + 1},
		} {
			ii, jj := nei[0], nei[1]
			if !ok(ii, jj) || seen[ii][jj] {
				continue
			}
			risk := x.val + uint16(grid[ii][jj])
			if risk >= dist[ii][jj] {
				continue
			}
			dist[ii][jj] = risk
			heap.Push(&h, pos{ii, jj, risk})
		}
	}
}

type pos struct {
	i, j int
	val  uint16
}
type minHeap []pos

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pos))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
