package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func toAnySlice(v any) []any {
	switch v.(type) {
	case []any:
		anySlice, success := v.([]any)
		if success {
			return anySlice
		} else {
			panic("oh no")
		}
	default:
		return []any{v}
	}
}

func inOrder(left, right any) int {
	leftFloat, isLeftFloat := left.(float64)
	rightFloat, isRightFloat := right.(float64)
	if isLeftFloat && isRightFloat {
		return int(rightFloat - leftFloat)
	}
	var leftSubSignal, rightSubSignal []any

	leftSubSignal = toAnySlice(left)
	rightSubSignal = toAnySlice(right)

	for i := range leftSubSignal {
		if len(rightSubSignal)-1 < i {
			return -1
		}
		if areInOrder := inOrder(leftSubSignal[i], rightSubSignal[i]); areInOrder != 0 {
			return areInOrder
		}
	}
	if len(leftSubSignal) == len(rightSubSignal) {
		return 0
	}
	return 1
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	pairs := strings.Split(string(input), "\n\n")
	total := 0
	allSignals := make([]any, 0, len(pairs)*2)
	for pairIdx, pair := range pairs {
		signals := strings.Split(pair, "\n")
		var left, right any

		json.Unmarshal([]byte(signals[0]), &left)
		allSignals = append(allSignals, left)
		json.Unmarshal([]byte(signals[1]), &right)
		allSignals = append(allSignals, right)

		if inOrder(left, right) >= 0 {
			total += pairIdx + 1
		}
		//fmt.Println(left, right.([]any), reflect.TypeOf(left))
		//os.Exit(0)
	}
	var parsedDivider any
	json.Unmarshal([]byte("[[2]]"), &parsedDivider)
	allSignals = append(allSignals, parsedDivider)
	json.Unmarshal([]byte("[[6]]"), &parsedDivider)
	allSignals = append(allSignals, parsedDivider)

	sort.Slice(allSignals, func(i, j int) bool { return inOrder(allSignals[i], allSignals[j]) > 0 })

	var firstDividerIdx, secondDividerIdx int
	for i, signal := range allSignals {
		serializedSignal, _ := json.Marshal(signal)

		switch string(serializedSignal) {
		case "[[2]]":
			firstDividerIdx = i + 1
		case "[[6]]":
			secondDividerIdx = i + 1
		}
	}

	fmt.Println(total)
	fmt.Println(firstDividerIdx * secondDividerIdx)
}
