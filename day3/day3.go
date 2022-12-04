package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input")

	inputString := string(input)
	bags := strings.Split(inputString, "\n")

	//Lowercase item types a through z have priorities 1 through 26.
	//Uppercase item types A through Z have priorities 27 through 52.

	//ascii
	// A-Z 65-90
	// a-z 97-122

	var total int
	for _, bag := range bags {
		contents := []rune(bag)

		firstHalf := contents[:len(contents)/2]
		secondHalf := contents[len(contents)/2:]

		total += intersect(firstHalf, secondHalf)
	}

	fmt.Println(total)

	var total2, groupCount int
	group := make([][]rune, 0, 3)
	for _, bag := range bags {

		group = append(group, []rune(bag))

		if len(group) == 3 {
			groupCount++
			total2 += tripleIntersect(group[0], group[1], group[2])
			group = nil
		}
	}
	fmt.Println(total2, groupCount)
}

func intersect(firstHalf, secondHalf []rune) int {
	for _, item1 := range firstHalf {
		for _, item2 := range secondHalf {
			if item1 == item2 {
				return getPriority(int(item1))
			}
		}
	}

	return 0
}

func tripleIntersect(a, b, c []rune) int {
	for _, item1 := range a {
		for _, item2 := range b {
			if item1 == item2 {
				for _, item3 := range c {
					if item1 == item3 {
						return getPriority(int(item3))
					}
				}
			}
		}
	}

	return 0
}

func getPriority(val int) int {
	if val > 90 {
		return val - 96
	}

	return val - 38
}
