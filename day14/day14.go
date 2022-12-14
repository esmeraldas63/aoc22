package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gridDimensions struct {
	startX int
	endX   int
	startY int
	endY   int
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	part := os.Args[2]
	rockPaths := strings.Split(string(input), "\n")
	grid := map[int]map[int]rune{}
	gDim := gridDimensions{999, 0, 0, 0}
	for _, rockPath := range rockPaths {
		prevX, prevY, currX, currY := -1, -1, 0, 0
		for _, coordinatesString := range strings.Split(rockPath, " -> ") {
			coordinates := strings.Split(coordinatesString, ",")
			currX, _ = strconv.Atoi(coordinates[0])
			currY, _ = strconv.Atoi(coordinates[1])

			if prevX != -1 {
				startY, endY := getEndpoints(prevY, currY)
				if endY > gDim.endY {
					gDim.endY = endY
				}
				startX, endX := getEndpoints(prevX, currX)
				if startX < gDim.startX {
					gDim.startX = startX
				}
				if endX > gDim.endX {
					gDim.endX = endX
				}
				for y := startY; y <= endY; y++ {
					if grid[y] == nil {
						grid[y] = map[int]rune{}
					}

					for x := startX; x <= endX; x++ {
						grid[y][x] = '#'
					}
				}
			}
			prevX, prevY = currX, currY
		}
	}

	if part == "2" {
		gDim.startX, gDim.endX, gDim.endY = 300, 700, gDim.endY+2
	}

	for y := gDim.startY; y <= gDim.endY; y++ {
		if grid[y] == nil {
			grid[y] = map[int]rune{}
		}
		for x := gDim.startX; x <= gDim.endX; x++ {
			if grid[y][x] == rune(0) {
				grid[y][x] = '.'

				if part == "2" && y == gDim.endY {
					grid[y][x] = '#'
				}
			}
		}
	}

	total := 0
	totallySettled := false
	for !totallySettled {
		y, x, settled := 0, 500, false
		for !settled {
			if grid[y+1][x] == '.' {
				y++
			} else if grid[y+1][x] == rune(0) {
				settled = true
				totallySettled = true
			} else if grid[y+1][x-1] == '.' {
				y++
				x--
			} else if grid[y+1][x-1] == rune(0) {
				settled = true
				totallySettled = true
			} else if grid[y+1][x+1] == '.' {
				y++
				x++
			} else if grid[y+1][x+1] == rune(0) {
				settled = true
				totallySettled = true
			} else {
				grid[y][x] = 'o'
				settled = true
				total++
			}

			if y == 0 && x == 500 {
				settled = true
				totallySettled = true
			}
		}
	}

	fmt.Println(total)
	printGrid(grid, gDim)
}

func printGrid(grid map[int]map[int]rune, gDim gridDimensions) {
	output, _ := os.Create("./aa")
	writer := bufio.NewWriter(output)

	for y := gDim.startY; y <= gDim.endY; y++ {
		line := ""
		for x := gDim.startX; x <= gDim.endX; x++ {
			line += string(grid[y][x])
		}

		writer.WriteString(line + "\n")
	}

	writer.Flush()
}

func getEndpoints(a, b int) (start, end int) {
	if b > a {
		return a, b
	}

	return b, a
}
