package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items          []*big.Int
	Operation      string
	TestDivisor    *big.Int
	IfTrue         int
	IfFalse        int
	InspectedItems int
}

func (monkey *Monkey) execOperation(modulo *big.Int) {
	operation := strings.Split(monkey.Operation, " ")

	var x, y *big.Int
	for i, item := range monkey.Items {
		if operation[0] == "old" {
			x = item
		} else {
			integer, _ := strconv.Atoi(operation[0])
			x = big.NewInt(int64(integer))
		}

		if operation[2] == "old" {
			y = item
		} else {
			integer, _ := strconv.Atoi(operation[2])
			y = big.NewInt(int64(integer))
		}

		switch operation[1] {
		case "+":
			monkey.Items[i].Add(x, y)
		case "-":
			monkey.Items[i].Sub(x, y)
		case "*":
			monkey.Items[i].Mul(x, y)
		case "/":
			monkey.Items[i].Div(x, y)
		default:
			panic(operation[1])
		}
		if modulo.BitLen() != 0 {
			monkey.Items[i].Mod(monkey.Items[i], modulo)
		} else {
			monkey.Items[i].Div(monkey.Items[i], big.NewInt(int64(3)))
		}
	}

}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	monkeyStrs := strings.Split(string(input), "\n\n")

	monkeys := make([]*Monkey, len(monkeyStrs))
	for i, monkeyStr := range monkeyStrs {
		monkey := Monkey{}
		monkeys[i] = &monkey

		monkeyProperties := strings.Split(monkeyStr, "\n")
		for _, item := range strings.Split(monkeyProperties[1][18:], ", ") {
			itemInt, _ := new(big.Int).SetString(item, 0)
			monkey.Items = append(monkey.Items, itemInt)
		}

		monkey.Operation = monkeyProperties[2][19:]
		monkey.TestDivisor, _ = new(big.Int).SetString(monkeyProperties[3][21:], 0)
		monkey.IfTrue, _ = strconv.Atoi(monkeyProperties[4][29:])
		monkey.IfFalse, _ = strconv.Atoi(monkeyProperties[5][30:])
	}

	var rounds int
	var modulo *big.Int
	if os.Args[2] == "1" {
		rounds = 20
		modulo = big.NewInt(0)
	} else {
		rounds = 10000
		modulo = big.NewInt(1)
		for _, monkey := range monkeys {
			modulo.Mul(modulo, monkey.TestDivisor)
		}
	}

	var modResult = new(big.Int)
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.InspectedItems += len(monkey.Items)
			monkey.execOperation(modulo)
			for _, item := range monkey.Items {
				if modResult.Mod(item, monkey.TestDivisor).BitLen() == 0 {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, item)
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, item)
				}
			}
			monkey.Items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectedItems > monkeys[j].InspectedItems
	})

	fmt.Println(monkeys[0].InspectedItems * monkeys[1].InspectedItems)
}
