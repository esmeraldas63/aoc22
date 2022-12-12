package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Y       int
	X       int
	Visited bool
	Val     rune
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(input), "\n")

	grid := make([][]Node, len(lines))
	//var start *Node
	var end, start *Node
	startingPoints := make([]*Node, 0, 50)
	for lineNum, line := range lines {
		grid[lineNum] = make([]Node, len(line))
		for colNum, val := range line {
			grid[lineNum][colNum] = Node{lineNum, colNum, false, val}
			if val == 'S' {
				start = &grid[lineNum][colNum]
				startingPoints = append(startingPoints, start)
				start.Val = 'a'
			}
			if val == 'a' {
				startingPoints = append(startingPoints, &grid[lineNum][colNum])
			}

			if val == 'E' {
				end = &grid[lineNum][colNum]
				end.Val = 'z'
			}
		}
	}

	var beingInspected *Node
	var neighbour Node
	newNeighbours := make([]Node, 0, 50)
	process := func() {
		if !beingInspected.Visited && beingInspected.Val-neighbour.Val <= 1 {
			newNeighbours = append(newNeighbours, *beingInspected)
			beingInspected.Visited = true

		}
	}
	stepCounts := make([]int, len(startingPoints))
	for i, startingPoint := range startingPoints {
		neighbours := []Node{*startingPoint}
		stepCount := 0

		for !end.Visited {
			stepCount++
			newNeighbours = nil

			if len(neighbours) == 0 {
				stepCount = 999
				break
			}

			for _, neighbour = range neighbours {
				neighbour.Visited = true
				if neighbour.X > 0 {
					beingInspected = &grid[neighbour.Y][neighbour.X-1]
					process()
				}

				if neighbour.X < len(grid[neighbour.Y])-1 {
					beingInspected = &grid[neighbour.Y][neighbour.X+1]
					process()
				}

				if neighbour.Y > 0 {
					beingInspected = &grid[neighbour.Y-1][neighbour.X]
					process()
				}

				if neighbour.Y < len(grid)-1 {
					beingInspected = &grid[neighbour.Y+1][neighbour.X]
					process()
				}

				neighbours = newNeighbours
			}
		}

		if startingPoint.X == start.X && startingPoint.Y == start.Y {
			fmt.Println(stepCount)
		}

		stepCounts[i] = stepCount

		for y, line := range grid {
			for x := range line {
				grid[y][x].Visited = false
			}
		}
	}

	sort.Ints(stepCounts)

	fmt.Println(stepCounts[0])
}
