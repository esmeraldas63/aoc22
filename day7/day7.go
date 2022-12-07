package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var dirSizes = make(map[string]int)
var dirMap = make(map[string][]string)

func main() {
	input, _ := os.ReadFile(os.Args[1])
	commands := strings.Split(string(input), "\n")[1:]
	sizeRe := regexp.MustCompile(`^\d+`)
	directories := make(map[string]string)
	pwd := ""
	directories[pwd] = pwd

	for _, command := range commands {
		if command[:4] == "$ ls" {
			continue
		}
		if command[:3] == "dir" {
			dirMap[pwd] = append(dirMap[pwd], pwd+"/"+command[4:])
			continue
		}

		if command[:4] == "$ cd" {
			dirName := command[5:]
			if dirName == ".." {
				pwd = pwd[:strings.LastIndex(pwd, "/")]
			} else {
				pwd += fmt.Sprintf("/%s", dirName)
				directories[pwd] = pwd
			}
			continue
		}

		sizeString := sizeRe.Find([]byte(command))
		fileSize, err := strconv.Atoi(string(sizeString))

		if err == nil {
			dirSizes[pwd] += fileSize
		} else {
			fmt.Println(err)
		}
	}

	totalSpaceUsed := getDirSize("")
	freeSpaceNeeded := 30000000 - (70000000 - totalSpaceUsed)
	toDelete := totalSpaceUsed
	total := 0
	for name, _ := range directories {
		dirSize := getDirSize(name)

		if freeSpaceNeeded <= dirSize && dirSize < toDelete {
			toDelete = dirSize
		}

		if dirSize <= 100000 {
			total += dirSize
		}
	}

	fmt.Println(total)
	fmt.Println(toDelete)
}

func getDirSize(dirName string) int {
	total := dirSizes[dirName]
	for _, subDir := range dirMap[dirName] {
		total += getDirSize(subDir)
	}

	return total
}
