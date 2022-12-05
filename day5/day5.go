package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var crates = [][]string{
	{"H", "L", "R", "F", "B", "C", "J", "M"},
	{"D", "C", "Z"},
	{"W", "G", "N", "C", "F", "J", "H"},
	{"B", "S", "T", "M", "D", "J", "P"},
	{"J", "R", "D", "C", "N"},
	{"Z", "G", "J", "P", "Q", "D", "L", "W"},
	{"H", "R", "F", "T", "Z", "P"},
	{"G", "M", "V", "L"},
	{"J", "R", "Q", "F", "P", "G", "B", "C"},
}

var crates2 = [][]string{
	{"H", "L", "R", "F", "B", "C", "J", "M"},
	{"D", "C", "Z"},
	{"W", "G", "N", "C", "F", "J", "H"},
	{"B", "S", "T", "M", "D", "J", "P"},
	{"J", "R", "D", "C", "N"},
	{"Z", "G", "J", "P", "Q", "D", "L", "W"},
	{"H", "R", "F", "T", "Z", "P"},
	{"G", "M", "V", "L"},
	{"J", "R", "Q", "F", "P", "G", "B", "C"},
}

func main() {
	actionsInput, _ := os.ReadFile("./actionInput")

	actionStrings := strings.Split(string(actionsInput), "\n")
	re := regexp.MustCompile(`\d+`)
	for _, actionString := range actionStrings {
		r := re.FindAll([]byte(actionString), -1)

		moveAmount, _ := strconv.Atoi(string(r[0]))
		moveFrom, _ := strconv.Atoi(string(r[1]))
		moveTo, _ := strconv.Atoi(string(r[2]))

		for i := 0; i < moveAmount; i++ {
			var x string

			stackToMoveFrom := &crates[moveFrom-1]
			stackToMoveTo := &crates[moveTo-1]
			x, *stackToMoveFrom = (*stackToMoveFrom)[0], (*stackToMoveFrom)[1:]
			*stackToMoveTo = append([]string{x}, *stackToMoveTo...)
		}

		stackToMoveFrom2 := &crates2[moveFrom-1]
		stackToMoveTo2 := &crates2[moveTo-1]
		var copyStack []string
		copyStack = append(copyStack, (*stackToMoveFrom2)[0:moveAmount]...)
		*stackToMoveTo2, *stackToMoveFrom2 = append(copyStack, *stackToMoveTo2...), (*stackToMoveFrom2)[moveAmount:]
	}

	var topCrates string
	for _, stack := range crates {
		topCrates += stack[0]
	}

	var topCrates2 string
	for _, stack := range crates2 {
		topCrates2 += stack[0]
	}

	fmt.Println(topCrates)
	fmt.Println(topCrates2)
}
