package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var snafuMap = map[rune]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

var inverseSnafuMap = map[int]rune{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	snafuNums := strings.Split(string(input), "\n")

	total := 0
	for _, snafuNum := range snafuNums {
		weight := 1
		decimal := 0
		for i := len(snafuNum) - 1; i >= 0; i-- {
			decimal += snafuMap[rune(snafuNum[i])] * weight
			weight *= 5
		}
		total += decimal
	}

	length := 27
	currIdx := 0
	remapped := make([]int, length)
	for i := int(math.Pow(5, float64(length-1))); i >= 1; i /= 5 {
		remapped[currIdx] = total / i
		if remapped[currIdx] > 2 {
			currOverflowIdx := currIdx
			for remapped[currOverflowIdx] > 2 {
				remapped[currOverflowIdx] -= 5
				remapped[currOverflowIdx-1] += 1
				currOverflowIdx--
			}
		}
		total %= i
		currIdx++
	}

	snafuTotal := ""
	for _, v := range remapped {
		snafuTotal += string(inverseSnafuMap[v])
	}

	fmt.Println(remapped, snafuTotal, total)
}
