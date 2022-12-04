package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input")

	inputString := string(input)
	elfStrs := strings.Split(inputString, "\n\n")

	elfTotals := make([]int, len(elfStrs))
	for i, valSliceString := range elfStrs {
		var total int
		for _, valString := range strings.Split(valSliceString, "\n") {
			valInt, _ := strconv.Atoi(valString)
			total += valInt
		}
		elfTotals[i] = total
	}

	sort.Slice(elfTotals, func(i, j int) bool {
		return elfTotals[i] > elfTotals[j]
	})

	fmt.Println(elfTotals[0])
	fmt.Println(elfTotals[0] + elfTotals[1] + elfTotals[2])
}
