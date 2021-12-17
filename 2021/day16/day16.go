package day16

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
)

//go:embed inputSmall.txt
var inputSmall string

//go:embed input.txt
var input string

func partA(in string) int64 {
	bin := hexToBin(in)
	for {
		if len(bin)%4 == 0 {
			break
		}
		bin = "0" + bin
	}
	_, sum := sumVal(bin, 0)
	return sum
}

func partB(in string) int64 {
	bin := hexToBin(in)
	for {
		if len(bin)%4 == 0 {
			break
		}
		bin = "0" + bin
	}
	_, val := calcVal(bin, 0)
	return val
}

func calcVal(bin string, pc int) (retPc int, ret int64) {
	pkgType := bin[pc+3 : pc+3+3]

	if pkgType == "100" {
		n := pc + 6
		var by string
		for {
			last := bin[n] == '0'
			by += bin[n+1 : n+5]
			n += 5
			if last {
				break
			}
		}
		pc = n
		return pc, binToDec(by)
	} else {
		lengthType := bin[pc+6]

		var subVals []int64

		if string(lengthType) == "0" {
			l := binToDec(bin[pc+7 : pc+7+15])
			pc = pc + 7 + 15
			stopAt := pc + int(l)

			for {
				nextpc, val := calcVal(bin, pc)
				pc = nextpc
				subVals = append(subVals, val)

				if pc >= stopAt {
					break
				}
			}

		} else {
			l := binToDec(bin[pc+7 : pc+7+11])
			pc = pc + 7 + 11

			for i := int64(0); i < l; i++ {
				nextpc, val := calcVal(bin, pc)
				pc = nextpc
				subVals = append(subVals, val)
			}
		}

		if pkgType == "000" {
			var s int64
			for _, v := range subVals {
				s += v
			}
			return pc, s
		}
		if pkgType == "001" {
			var s int64 = 1
			for _, v := range subVals {
				s *= v
			}
			return pc, s
		}
		if pkgType == "010" {
			var s int64 = math.MaxInt64
			for _, v := range subVals {
				s = min(s, v)
			}
			return pc, s
		}
		if pkgType == "011" {
			var s int64
			for _, v := range subVals {
				s = max(s, v)
			}
			return pc, s
		}
		if pkgType == "101" {
			var s int64
			if subVals[0] > subVals[1] {
				s = 1
			}
			return pc, s
		}
		if pkgType == "110" {
			var s int64
			if subVals[0] < subVals[1] {
				s = 1
			}
			return pc, s
		}
		if pkgType == "111" {
			var s int64
			if subVals[0] == subVals[1] {
				s = 1
			}
			return pc, s
		}
	}
	return
}

func sumVal(bin string, pc int) (retPc int, sumVersion int64) {
	version := binToDec(bin[pc+0 : pc+0+3])
	pkgType := bin[pc+3 : pc+3+3]
	if pkgType == "100" {
		n := pc + 6
		var by string
		for {
			last := bin[n] == '0'
			by += bin[n+1 : n+5]
			n += 5
			if last {
				break
			}
		}
		pc = n
		return pc, version
	} else {
		lengthType := bin[pc+6]
		if string(lengthType) == "0" {
			l := binToDec(bin[pc+7 : pc+7+15])
			pc = pc + 7 + 15
			stopAt := pc + int(l)
			for {
				nextpc, addV := sumVal(bin, pc)
				version += addV
				pc = nextpc
				if pc >= stopAt {
					break
				}
			}
		} else {
			l := binToDec(bin[pc+7 : pc+7+11])
			pc = pc + 7 + 11
			for i := int64(0); i < l; i++ {
				nextpc, addV := sumVal(bin, pc)
				version += addV
				pc = nextpc
			}
		}
		return pc, version
	}
}

func hexToBin(in string) string {
	decoded, err := hex.DecodeString(in)
	if err != nil {
		log.Fatal(err)
	}
	var res string
	for _, d := range decoded {
		res += fmt.Sprintf("%08b", d)
	}
	return res

}

func binToDec(bin string) int64 {
	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return num

}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y

}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y

}
