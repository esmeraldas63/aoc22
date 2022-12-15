package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type gridDimensions struct {
	startX int
	endX   int
	startY int
	endY   int
}

func (gDim *gridDimensions) update(x, y int) {
	if y > gDim.endY {
		gDim.endY = y
	}
	if y < gDim.startY {
		gDim.startY = y
	}

	if x > gDim.endX {
		gDim.endX = x
	}
	if x < gDim.startX {
		gDim.startX = x
	}
}

type pair struct {
	sX       int
	sY       int
	bX       int
	bY       int
	distance int
}

func (p *pair) getDistance() int {
	if p.distance == 0 {
		p.distance = manhattanDistance(p.sX, p.bX, p.sY, p.bY)
	}

	return p.distance
}

func (p *pair) inSensorProximity(x, y int) bool {
	if manhattanDistance(p.sX, x, p.sY, y) <= p.getDistance() {
		return true
	}

	return false
}

var pairs []pair

func main() {
	input, _ := os.ReadFile(os.Args[1])
	part := os.Args[2]
	inputLines := strings.Split(string(input), "\n")
	gDim := gridDimensions{999, 0, 999, 0}
	sensorXRe := regexp.MustCompile(`Sensor at x=(\-?\d+)`)
	sensorYRe := regexp.MustCompile(`Sensor at x=\-?\d+, y=(\-?\d+)`)
	beaconXRe := regexp.MustCompile(`beacon is at x=(\-?\d+)`)
	beaconYRe := regexp.MustCompile(`beacon is at x=\-?\d+, y=(\-?\d+)`)
	grid := map[int]map[int]rune{}

	var matches [][]byte
	pairs = make([]pair, len(inputLines))
	for i, inputLine := range inputLines {
		matches = sensorXRe.FindSubmatch([]byte(inputLine))
		pairs[i].sX, _ = strconv.Atoi(string(matches[1]))
		matches = sensorYRe.FindSubmatch([]byte(inputLine))
		pairs[i].sY, _ = strconv.Atoi(string(matches[1]))
		if grid[pairs[i].sY] == nil {
			grid[pairs[i].sY] = map[int]rune{}
		}
		grid[pairs[i].sY][pairs[i].sX] = 'S'
		matches = beaconXRe.FindSubmatch([]byte(inputLine))
		pairs[i].bX, _ = strconv.Atoi(string(matches[1]))
		matches = beaconYRe.FindSubmatch([]byte(inputLine))
		pairs[i].bY, _ = strconv.Atoi(string(matches[1]))
		if grid[pairs[i].bY] == nil {
			grid[pairs[i].bY] = map[int]rune{}
		}
		grid[pairs[i].bY][pairs[i].bX] = 'B'

		gDim.update(pairs[i].sY, pairs[i].sX)
		gDim.update(pairs[i].bY, pairs[i].bX)
	}

	if part == "1" {
		lineToInspect := 2000000
		total := 0
		for _, pair := range pairs {
			if grid[lineToInspect] == nil {
				grid[lineToInspect] = map[int]rune{}
			}
			for x := pair.sX - pair.getDistance(); x <= pair.sX+pair.getDistance(); x++ {
				if grid[lineToInspect][x] == rune(0) && pair.inSensorProximity(x, lineToInspect) {
					grid[lineToInspect][x] = '#'
					total++
				}
			}
		}
		fmt.Println(total)
	}

	if part == "2" {
		for _, pair := range pairs {
			count, direction := 0, 1
			fmt.Println(pair.getDistance())
			for y := pair.sY - pair.getDistance() - 1; y <= pair.sY+pair.getDistance()+1; y++ {
				checkPartTwo(pair.sX-count, y)
				checkPartTwo(pair.sX+count, y)
				if y == pair.sY {
					direction = -1
				}
				count += direction
			}
		}
	}
}

func checkPartTwo(x, y int) {
	if x < 0 || x > 4000000 {
		return
	}
	if y < 0 || y > 4000000 {
		return
	}
	for _, pair := range pairs {
		if pair.inSensorProximity(x, y) {
			return
		}
	}

	fmt.Println(x*4000000 + y)
	os.Exit(0)
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func manhattanDistance(x1, x2, y1, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
