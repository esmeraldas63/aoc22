package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile(os.Args[1])
	commands := strings.Split(string(input), "\n")

	x := 1
	cycle := 1

	screen := make([][]bool, 6)
	for i := 0; i < len(screen); i++ {
		screen[i] = make([]bool, 40)
	}

	checkQueue := []int{60, 100, 140, 180, 220}
	nextCheck := 20
	total := 0
	check := func() {
		trueCycle := cycle - 1
		if abs((trueCycle%40)-x) <= 1 {
			screen[trueCycle/40][trueCycle%40] = true
		}

		if cycle%nextCheck == 0 {
			total += x * cycle

			if len(checkQueue) > 0 {
				nextCheck, checkQueue = checkQueue[0], checkQueue[1:]
			} else {
				fmt.Println(total)
			}
		}
	}

	check()
	for _, command := range commands {
		switch command[:4] {
		case "noop":
			cycle++
			check()
		case "addx":
			cycle++
			check()
			cycle++
			increaseBy, _ := strconv.Atoi(command[5:])
			x += increaseBy
			check()
		}
	}

	paintScreen(screen)
}

func paintScreen(screen [][]bool) {
	for y := 0; y < 6; y++ {
		line := ""
		for x := 0; x < 40; x++ {
			if screen[y][x] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x - 0
}
