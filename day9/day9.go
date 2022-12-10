package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	moves := strings.Split(string(input), "\n")

	visited := make(map[int]map[int]bool)

	ropeCount := 10
	knots := make([]Position, ropeCount)
	tail := &knots[ropeCount-1]
	head := &knots[0]
	leftTop := Position{0, 0}
	rightBottom := Position{0, 0}
	visited[0] = make(map[int]bool)
	for _, move := range moves {
		direction := move[0]
		moveCount, _ := strconv.Atoi(move[2:])
		for i := 0; i < moveCount; i++ {
			switch direction {
			case 'U':
				head.Y++
			case 'D':
				head.Y--
			case 'R':
				head.X++
			case 'L':
				head.X--
			}
			for i := 1; i < ropeCount; i++ {
				if knots[i].Y+1 < knots[i-1].Y && knots[i].X+1 < knots[i-1].X {
					knots[i].X, knots[i].Y = knots[i-1].X-1, knots[i-1].Y-1
				} else if knots[i].Y+1 < knots[i-1].Y && knots[i].X-1 > knots[i-1].X {
					knots[i].X, knots[i].Y = knots[i-1].X+1, knots[i-1].Y-1
				} else if knots[i].Y-1 > knots[i-1].Y && knots[i].X+1 < knots[i-1].X {
					knots[i].X, knots[i].Y = knots[i-1].X-1, knots[i-1].Y+1
				} else if knots[i].Y-1 > knots[i-1].Y && knots[i].X-1 > knots[i-1].X {
					knots[i].X, knots[i].Y = knots[i-1].X+1, knots[i-1].Y+1
				} else if knots[i].Y+1 < knots[i-1].Y {
					knots[i].X, knots[i].Y = knots[i-1].X, knots[i-1].Y-1
				} else if knots[i].Y-1 > knots[i-1].Y {
					knots[i].X, knots[i].Y = knots[i-1].X, knots[i-1].Y+1
				} else if knots[i].X+1 < knots[i-1].X {
					knots[i].X, knots[i].Y = knots[i-1].X-1, knots[i-1].Y
				} else if knots[i].X-1 > knots[i-1].X {
					knots[i].X, knots[i].Y = knots[i-1].X+1, knots[i-1].Y
				}
			}
			if visited[tail.Y] == nil {
				visited[tail.Y] = make(map[int]bool)
			}
			visited[tail.Y][tail.X] = true

			if leftTop.Y < tail.Y {
				leftTop.Y = tail.Y
			} else if rightBottom.Y > tail.Y {
				rightBottom.Y = tail.Y
			}

			if leftTop.X > tail.X {
				leftTop.X = tail.X
			} else if rightBottom.X < tail.X {
				rightBottom.X = tail.X
			}

		}
	}

	visitedCount := 0
	for _, v := range visited {
		visitedCount += len(v)
	}

	printGrid(leftTop, rightBottom, visited)

	fmt.Println(visitedCount)
}

func printGrid(leftTop Position, rightBottom Position, visited map[int]map[int]bool) {
	output, _ := os.Create("./part2")
	writer := bufio.NewWriter(output)

	for y := leftTop.Y; y >= rightBottom.Y; y-- {
		line := ""
		for x := leftTop.X; x <= rightBottom.X; x++ {
			if visited[y][x] {
				line += "#"
			} else {
				line += "."
			}
		}
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
}
