package day12

import (
	_ "embed"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

type caves map[string]map[string]bool

func partA(input string) int {
	c := caves{}

	var ans [][]string
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, "-"))
	}

	for _, pair := range ans {
		if c[pair[0]] == nil {
			c[pair[0]] = map[string]bool{}
		}
		if c[pair[1]] == nil {
			c[pair[1]] = map[string]bool{}
		}
		c[pair[0]][pair[1]] = true
		c[pair[1]][pair[0]] = true
	}

	return c.countPaths("start", map[string]bool{"start": true})
}

func partB(input string) int {
	c := caves{}

	var ans [][]string
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, "-"))
	}

	for _, pair := range ans {
		if c[pair[0]] == nil {
			c[pair[0]] = map[string]bool{}
		}
		if c[pair[1]] == nil {
			c[pair[1]] = map[string]bool{}
		}
		c[pair[0]][pair[1]] = true
		c[pair[1]][pair[0]] = true
	}

	return c.countPaths2("start", map[string]int{"start": 5}, false)
}

func (c caves) countPaths(current string, visited map[string]bool) int {
	if current == "end" {
		return 1
	}

	var pathsToEnd int

	for visitable := range c[current] {
		if visited[visitable] && strings.ToUpper(visitable) != visitable {
			continue
		}
		visited[current] = true
		pathsToEnd += c.countPaths(visitable, visited)

		visited[visitable] = false
	}

	return pathsToEnd
}

func (c caves) countPaths2(current string, visited map[string]int, doubleUsed bool) int {
	if current == "end" {
		return 1
	}

	visited[current]++

	var pathsToEnd int

	for visitable := range c[current] {
		if visitable == "start" {
			continue
		}

		if strings.ToUpper(visitable) != visitable && visited[visitable] > 0 {
			if doubleUsed {
				continue
			} else {
				doubleUsed = true
			}
		}

		pathsToEnd += c.countPaths2(visitable, visited, doubleUsed)

		visited[visitable]--
		if strings.ToUpper(visitable) != visitable && visited[visitable] == 1 {
			doubleUsed = false
		}
	}

	return pathsToEnd
}
