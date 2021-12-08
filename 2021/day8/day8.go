package day8

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(input string) int {
	counter := 0
	wordLen := 0

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " | ")
		for i := range parts[1] {
			if parts[1][i] == ' ' {
				// 2 for digit 1
				// 4 for digit 4
				// 3 for digit 7
				// 7 for digit 8
				if wordLen == 2 || wordLen == 4 || wordLen == 3 || wordLen == 7 {
					counter++
				}
				wordLen = 0
				continue
			}

			wordLen++
		}

		if wordLen == 2 || wordLen == 4 || wordLen == 3 || wordLen == 7 {
			counter++
		}

		wordLen = 0
	}

	return counter
}

func partB(input string) int {
	lines := strings.Split(input, "\n")
	digits := [10]string{
		"cagedb",  // 0
		"ab",      // 1
		"gcdfa",   // 2
		"fbcad",   // 3
		"eafb",    // 4
		"cdfbe",   // 5
		"cdfgeb",  // 6
		"dab",     // 7
		"acedgfb", // 8
		"cefabd",  // 9
	}

	var counter int

	for _, line := range lines {
		p := strings.Split(line, " | ")
		inputs := strings.Split(p[0], " ")
		outputs := strings.Split(p[1], " ")
		wires := [10]string{}

		for _, v := range inputs {
			l := len(v)
			if l == 2 {
				wires[1] = v
			} else if l == 3 {
				wires[7] = v
			} else if l == 4 {
				wires[4] = v
			} else if l == 7 {
				wires[8] = v
			}
		}

		for i := 0; i < 10; i++ {
			for _, v := range inputs {
				if wires[2] == "" && len(v) == 5 && overlaps(v, wires[7]) == 2 && overlaps(v, wires[1]) == 1 && overlaps(v, wires[3]) == 4 && overlaps(v, wires[4]) == 2 {
					wires[2] = v
				}

				if wires[3] == "" && len(v) == 5 && overlaps(v, wires[1]) == 2 {
					wires[3] = v
				}

				if wires[9] == "" && len(v) == 6 && overlaps(v, wires[7]) == 3 && overlaps(v, wires[0]) == 5 {
					wires[9] = v
				}

				if wires[5] == "" && len(v) == 5 && overlaps(v, wires[2]) == 3 && overlaps(v, wires[4]) == 3 {
					wires[5] = v
				}

				if wires[6] == "" && len(v) == 6 && overlaps(v, wires[1]) == 1 && overlaps(v, wires[2]) == 4 {
					wires[6] = v
				}

				if wires[0] == "" && len(v) == 6 && overlaps(v, wires[1]) == 2 && overlaps(v, wires[2]) == 4 && overlaps(v, wires[3]) == 4 {
					wires[0] = v
				}

			}

		}

		var o string

		for k, v := range wires {
			wires[k] = sortString(v)
		}

		for k, v := range digits {
			digits[k] = sortString(v)
		}

		for _, v := range outputs {
			v = sortString(v)
			for k, w := range wires {
				if w == v {
					o += fmt.Sprintf("%d", k)
				}
			}
		}

		counter += strToInt(o)
	}

	return counter
}

func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func overlaps(a, b string) (count int) {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			count++
		}
	}
	return
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")

}
