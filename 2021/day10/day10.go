package day10

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(input string) int {
	score := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		score += calcLineScore(line)
	}

	return score
}

func partB(input string) int {
	scores := []int{}

	i1 := 0
	for i := range input {
		if input[i] == '\n' {
			s := calcLineIncompleteScore(input[i1:i])
			if s != 0 {
				scores = append(scores, s)
			}

			i1 = i + 1
		}
	}

	sort.Ints(scores)

	return scores[int(len(scores)/2)]
}

func calcLineIncompleteScore(line string) int {
	scores := map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	openChars := make([]byte, len(line))
	pos := -1

	for i := range line {
		switch line[i] {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '<':
			pos++
			openChars[pos] = line[i]
			// openChars = append(openChars, line[i])
			continue
		}

		lastChar := openChars[pos]

		switch line[i] {
		case ')':
			if lastChar != '(' {
				return 0
			}
			pos--
			// openChars = openChars[:len(openChars)-1]
		case ']':
			if lastChar != '[' {
				return 0
			}
			// openChars = openChars[:len(openChars)-1]
			pos--
		case '}':
			if lastChar != '{' {
				return 0
			}
			// openChars = openChars[:len(openChars)-1]
			pos--
		case '>':
			if lastChar != '<' {
				return 0
			}
			// openChars = openChars[:len(openChars)-1]
			pos--
		}
	}

	score := 0

	for i := len(openChars) - 1; i >= 0; i-- {
		score = score*5 + scores[openChars[i]]
	}

	return score
}

func calcLineScore(line string) int {
	scores := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	openChars := make([]byte, 0)

	for i := range line {
		switch line[i] {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '<':
			openChars = append(openChars, line[i])
			continue
		}

		lastChar := openChars[len(openChars)-1]

		switch line[i] {
		case ')':
			if lastChar != '(' {
				return scores[line[i]]
			}
			openChars = openChars[:len(openChars)-1]
		case ']':
			if lastChar != '[' {
				return scores[line[i]]
			}
			openChars = openChars[:len(openChars)-1]
		case '}':
			if lastChar != '{' {
				return scores[line[i]]
			}
			openChars = openChars[:len(openChars)-1]
		case '>':
			if lastChar != '<' {
				return scores[line[i]]
			}
			openChars = openChars[:len(openChars)-1]
		}
	}

	return 0
}
