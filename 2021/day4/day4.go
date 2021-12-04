package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day4.txt
var input string

type board [5][5]int

func (b *board) sum() int {
	sum := 0
	for x := range b {
		for y := range b[x] {
			if b[x][y] != -1 {
				sum += b[x][y]
			}
		}
	}

	return sum
}

func (b *board) markNumber(n int) {
	for x := range b {
		for y := range b[x] {
			if b[x][y] == n {
				b[x][y] = -1
			}
		}
	}
}

func (b *board) hasBingo() bool {
	for x := range b {
		c := 0
		for y := range b[x] {
			if b[x][y] == -1 {
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
			if b[x][y] == -1 {
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
	partA()
	partB()
}

func partA() {
	input := strings.NewReader(input)
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	numbers := stringsToInts(strings.Split(scanner.Text(), ","))
	boards := readBoards(scanner)

	for _, n := range numbers {
		for i := range boards {
			boards[i].markNumber(n)
			if boards[i].hasBingo() {
				fmt.Printf("BINGO! : %v\n, sum: %d\n", boards[i], n*boards[i].sum())
				return
			}
		}
	}
}

func partB() {
	input := strings.NewReader(input)
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	numbers := stringsToInts(strings.Split(scanner.Text(), ","))
	boards := readBoards(scanner)

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
					fmt.Printf("Last BINGO! : %v\n, sum: %d\n", boards[i], n*boards[i].sum())
				}
			}
		}
	}
}

func readBoards(scanner *bufio.Scanner) []board {
	boards := []board{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		b := board{}
		b[0] = stringsToInts5(strings.Split(line, " "))

		for i := 1; i < 5; i++ {
			scanner.Scan()
			line := scanner.Text()
			b[i] = stringsToInts5(strings.Split(line, " "))
		}

		boards = append(boards, b)
	}

	return boards
}

func stringsToInts5(s []string) [5]int {
	ints := [5]int{}
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}
		mes, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}

		ints[index] = mes
		index++
	}
	return ints
}

func stringsToInts(s []string) []int {
	ints := []int{}
	for i := 0; i < len(s); i++ {
		mes, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, mes)
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
