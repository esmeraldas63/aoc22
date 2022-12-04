package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input")

	inputString := string(input)
	combinationStrs := strings.Split(inputString, "\n")

	//A for Rock, B for Paper, and C for Scissors. The second column--
	//1 for Rock, 2 for Paper, and 3 for Scissors
	//X for Rock, Y for Paper, and Z for Scissors
	//X means you need to lose, Y means you need to end the round in a draw, and Z means you
	winMap := map[string]int{"A": 2, "B": 3, "C": 1}
	drawMap := map[string]int{"A": 1, "B": 2, "C": 3}
	loseMap := map[string]int{"A": 3, "B": 1, "C": 2}
	//winMap := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	//equalMap := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	valueMap := map[string]int{"X": 0, "Y": 3, "Z": 6}
	//valueMap2 := map[string]int{"A": 2, "B": 3, "C": 1}

	var total int
	for _, v := range combinationStrs {
		comb := strings.Split(v, " ")

		//if winMap[comb[0]] == comb[1] {
		//	total += 6
		//}
		//
		//if equalMap[comb[0]] == comb[1] {
		//	total += 3
		//}
		currVal := valueMap[comb[1]]
		total += currVal

		if currVal == 6 {
			total += winMap[comb[0]]
		}

		if currVal == 3 {
			total += drawMap[comb[0]]
		}

		if currVal == 0 {
			total += loseMap[comb[0]]
		}

	}

	fmt.Println(total)
}
