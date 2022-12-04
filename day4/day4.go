package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input")

	pairs := strings.Split(string(input), "\n")

	var totalCompleteOverlap int
	var totalOverlap int
	//pairs := make([]int, len(pairsStr))
	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")
		elf1 := strings.Split(elfs[0], "-")
		elf2 := strings.Split(elfs[1], "-")
		if checkCompleteOverlap(stringsToInts(elf1), stringsToInts(elf2)) {
			totalCompleteOverlap++
		}

		if checkOverlap(stringsToInts(elf1), stringsToInts(elf2)) {
			totalOverlap++
		}
	}

	fmt.Println(totalCompleteOverlap)
	fmt.Println(totalOverlap)
}

func checkCompleteOverlap(range1, range2 []int) bool {
	return (range1[0] >= range2[0] && range1[1] <= range2[1]) || (range1[0] <= range2[0] && range1[1] >= range2[1])
}

func checkOverlap(range1, range2 []int) bool {
	return range1[0] <= range2[1] && range2[0] <= range1[1]
}

func stringsToInts(strings []string) []int {
	var ints = make([]int, len(strings))

	for i, v := range strings {
		int, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		ints[i] = int
	}

	return ints
}
