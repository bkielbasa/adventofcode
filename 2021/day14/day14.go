package day14

import (
	_ "embed"
	"strings"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partB(in string, steps int) int {
	lines := strings.Split(in, "\n")
	polimer := lines[0]
	polymerMap := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		l := strings.Split(lines[i], " -> ")
		polymerMap[l[0]] = l[1]
	}

	charMap := make(map[string]int)

	for i := 0; i < len(polimer)-1; i++ {
		charMap[polimer[i:i+2]] += 1
	}
	charMap[polimer[len(polimer)-1:]] += 1

	for i := 0; i < steps; i++ {
		newCharMap := make(map[string]int)

		for k, v := range charMap {
			if polymerMap[k] != "" {
				newCharMap[k[0:1]+polymerMap[k]] += v
				newCharMap[polymerMap[k]+k[1:2]] += v
			} else {
				newCharMap[k] += v
			}
		}
		charMap = newCharMap
	}

	freq := make(map[string]int)
	for k, v := range charMap {
		freq[k[0:1]] += v
	}

	return sumOccurrencesInMap(freq)
}

func partA(in string, steps int) int {
	lines := strings.Split(in, "\n")
	i := 2
	polimer := lines[0]
	ip := map[string]string{}

	for ; i < len(lines); i++ {
		l := strings.Split(lines[i], " -> ")
		ip[l[0]] = l[1]
	}

	for i = 0; i < steps; i++ {
		newPolimer := strings.Builder{}
		for index := 0; index < len(polimer)-1; index++ {
			pair := polimer[index : index+2]
			if v, ok := ip[pair]; ok {
				newPolimer.WriteByte(pair[0])
				newPolimer.WriteString(v)
			} else {
				newPolimer.WriteByte(pair[0])
			}
		}

		newPolimer.WriteByte(polimer[len(polimer)-1])

		polimer = newPolimer.String()
	}

	return sumOccurrences(polimer)
}

func sumOccurrencesInMap(m map[string]int) int {
	least := 999999999999
	most := 0

	for _, v := range m {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}

func sumOccurrences(str string) int {
	m := map[rune]int{}
	for _, s := range str {
		m[s]++
	}

	least := 999999999
	most := 0

	for _, v := range m {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}
