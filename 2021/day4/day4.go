package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day4.txt
var input string

type board [5][5]string

func (b *board) sum() int {
	sum := 0
	for x := range b {
		for y := range b[x] {
			if b[x][y] != "-1" {
				s, _ := strconv.Atoi(b[x][y])
				sum += s
			}
		}
	}

	return sum
}

func (b *board) markNumber(n string) {
	for x := range b {
		for y := range b[x] {
			if b[x][y] == n {
				b[x][y] = "-1"
			}
		}
	}
}

func (b *board) hasBingo() bool {
	for x := range b {
		c := 0
		for y := range b[x] {
			if b[x][y] == "-1" {
				c++
			}
		}

		if c == 5 {
			return true
		}
	}

	for y := 0; y < 5; y++ {
		c := 0

		for x := range b {
			if b[x][y] == "-1" {
				c++
			}
		}

		if c == 5 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(partA())
	fmt.Println(partB())
}

func partA() int {
	noOfBoards := (len(strings.Split(input, "\n")) - 1) / 6
	input := strings.NewReader(input)
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	boards := readBoards(scanner, noOfBoards)

	for _, n := range numbers {
		for i := range boards {
			boards[i].markNumber(n)
			if boards[i].hasBingo() {
				d, _ := strconv.Atoi(n)
				return d * boards[i].sum()
			}
		}
	}

	return 0
}

func partB() int {
	noOfBoards := (len(strings.Split(input, "\n")) - 1) / 6
	input := strings.NewReader(input)
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	boards := readBoards(scanner, noOfBoards)

	alreadyWon := []int{}
	allBoards := len(boards)
	for _, n := range numbers {
		for i := range boards {
			if intsInSlice(i, alreadyWon) {
				continue
			}

			boards[i].markNumber(n)
			if boards[i].hasBingo() {
				alreadyWon = append(alreadyWon, i)
				if len(alreadyWon) == allBoards {
					s, _ := strconv.Atoi(n)
					return s * boards[i].sum()
				}
			}
		}
	}

	return 0
}

func readBoards(scanner *bufio.Scanner, max int) []board {
	boards := make([]board, max)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		b := board{}
		b[0] = strings5(strings.Split(line, " "))

		for i := 1; i < 5; i++ {
			scanner.Scan()
			line := scanner.Text()
			b[i] = strings5(strings.Split(line, " "))
		}

		boards = append(boards, b)
	}

	return boards
}

func strings5(s []string) [5]string {
	ints := [5]string{}
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}

		ints[index] = s[i]
		index++
	}
	return ints
}

func intsInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
