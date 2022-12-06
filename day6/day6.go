package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("./input")
	inputString := string(input)

	singalLenght := 14
	signal := make([]rune, 0, singalLenght)
	for i, inputRune := range inputString {
		if len(signal) < singalLenght {
			signal = append(signal, inputRune)
		} else {
			//fmt.Println(signal)
			if isUnique(signal) {
				fmt.Println(i)
				os.Exit(0)
			}
			signal = append(signal[1:], inputRune)
		}
	}
}

func isUnique(slice []rune) bool {
	sliceLen := len(slice)
	for i, v := range slice {
		for j := i + 1; j < sliceLen; j++ {
			if v == slice[j] {
				return false
			}
		}
	}

	return true
}
