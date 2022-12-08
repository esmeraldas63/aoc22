package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(input), "\n")

	highestScenicScore := 0
	columnCount := len(lines[0])
	lineCount := len(lines)
	for treeLine, line := range lines {
		for treeCol, tree := range line {
			treeHeight, _ := strconv.Atoi(string(tree))

			c1, c2, c3, c4 := 0, 0, 0, 0
			if treeLine+1 != lineCount {
				for line := treeLine + 1; line < lineCount; line++ {
					nHeight, _ := strconv.Atoi(string(lines[line][treeCol]))
					c1++
					if treeHeight <= nHeight {
						break
					}
				}
			}

			if treeLine != 0 {
				for line := treeLine - 1; line >= 0; line-- {
					nHeight, _ := strconv.Atoi(string(lines[line][treeCol]))
					c2++
					if treeHeight <= nHeight {
						break
					}
				}
			}

			if treeCol != 0 {
				for col := treeCol - 1; col >= 0; col-- {
					nHeight, _ := strconv.Atoi(string(lines[treeLine][col]))
					c3++
					if treeHeight <= nHeight {
						break
					}
				}
			}

			if treeCol+1 != columnCount {
				for col := treeCol + 1; col < columnCount; col++ {
					nHeight, _ := strconv.Atoi(string(lines[treeLine][col]))
					c4++
					if treeHeight <= nHeight {
						break
					}
				}
			}

			if c1 == 0 {
				c1 = 1
			}
			if c2 == 0 {
				c2 = 1
			}
			if c3 == 0 {
				c3 = 1
			}
			if c4 == 0 {
				c4 = 1
			}
			tss := c1 * c2 * c3 * c4

			if tss > highestScenicScore {
				highestScenicScore = tss
			}
		}
	}

	var tallestTree int
	visibleGrid := make(map[int]map[int]bool)
	for lineNum, line := range lines {
		tallestTree = -1
		for columnNum, tree := range line {
			treeHeight, _ := strconv.Atoi(string(tree))
			if treeHeight > tallestTree {
				if visibleGrid[lineNum] == nil {
					visibleGrid[lineNum] = make(map[int]bool)
				}
				visibleGrid[lineNum][columnNum] = true
				tallestTree = treeHeight
			}
		}
	}

	for lineNum, line := range lines {
		tallestTree = -1
		for columnNum := len(line) - 1; columnNum >= 0; columnNum-- {
			treeHeight, _ := strconv.Atoi(string(line[columnNum]))
			if treeHeight > tallestTree {
				if visibleGrid[lineNum] == nil {
					visibleGrid[lineNum] = make(map[int]bool)
				}
				visibleGrid[lineNum][columnNum] = true
				tallestTree = treeHeight
			}
		}
	}

	for columnNum := 0; columnNum < columnCount; columnNum++ {
		tallestTree = -1
		for lineNum := 0; lineNum < len(lines); lineNum++ {
			treeHeight, _ := strconv.Atoi(string(lines[lineNum][columnNum]))
			if treeHeight > tallestTree {
				if visibleGrid[lineNum] == nil {
					visibleGrid[lineNum] = make(map[int]bool)
				}

				visibleGrid[lineNum][columnNum] = true
				tallestTree = treeHeight
			}
		}
	}

	for columnNum := 0; columnNum < columnCount; columnNum++ {
		tallestTree = -1
		for lineNum := len(lines) - 1; lineNum >= 0; lineNum-- {
			treeHeight, _ := strconv.Atoi(string(lines[lineNum][columnNum]))
			if treeHeight > tallestTree {
				if visibleGrid[lineNum] == nil {
					visibleGrid[lineNum] = make(map[int]bool)
				}

				visibleGrid[lineNum][columnNum] = true
				tallestTree = treeHeight
			}

		}
	}

	total := 0
	for _, visibleLine := range visibleGrid {
		total += len(visibleLine)
	}

	fmt.Println(total)
	fmt.Println(highestScenicScore)
}
