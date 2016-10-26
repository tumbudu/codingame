//Code of the Rings


package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
type stone struct {
	value uint8
	// op        uint8
	valueDist int
	weight    int
}

func valDiff(target, current uint8) int {
	t := m[string(target)]
	c := m[string(current)]

	diff := t - c
	// fmt.Println("target values", t, "and ", c, "is", diff)
	if diff < 0 {
		if diff < -14 {
			diff = 27 - (diff * -1)
		}
		// fmt.Println("Differnece between", string(target), "and ", string(current), "is", diff)
		return diff
	}

	if diff > 0 {
		if diff > 14 {
			diff = 27 - diff
			// fmt.Println("Differnece between", string(target), "and ", string(current), "is", diff)
			return diff * -1
		} else {
			// fmt.Println("Differnece between", string(target), "and ", string(current), "is", diff)
			return diff
		}
	}
	// fmt.Println("Differnece between", string(target), "and ", string(current), "is", diff)
	return diff
}

var m map[string]int
var stoneList []stone

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	magicPhrase := scanner.Text()

	stoneList = make([]stone, 30)

	for i := 0; i < 30; i++ {
		stoneList[i].value = ' '
		stoneList[i].weight = 1000
	}
	currentIndex := 0
	initMap()
	// fmt.Fprintln(os.Stderr, "String is", magicPhrase," With length", len(magicPhrase))
	// var X uint8
	// X = ' '
	for j := 0; j < len(magicPhrase); j++ {
		c := magicPhrase[j]
		// fmt.Fprintln(os.Stderr, string(c))

		for i := 0; i < 30; i++ {
			stoneList[i].valueDist = valDiff(c, stoneList[i].value)

			stoneDiff := currentIndex - i
			if stoneDiff < 0 {
				stoneDiff *= -1
			}
			if stoneDiff > 15 {
				stoneDiff = 30 - stoneDiff
			}
			if stoneList[i].valueDist < 0 {
				stoneList[i].weight = stoneList[i].valueDist * -1
			} else {
				stoneList[i].weight = stoneList[i].valueDist
			}

			stoneList[i].weight += stoneDiff
		}

		lowIndex := currentIndex
		lowWt := stoneList[currentIndex].weight

		for i := 0; i < 30; i++ {
			if lowWt > stoneList[i].weight {
				lowWt = stoneList[i].weight
				lowIndex = i
			}
		}

		// fmt.Println("Operate on stone(Low Weight) ", lowIndex)
		operate(lowIndex, currentIndex, c)
		currentIndex = lowIndex

		// for i := 0; i < 30; i++ {
		// 	fmt.Println(stoneList[i])
		// }
		// fmt.Println()
		// fmt.Println()
		// fmt.Println()
	}
	fmt.Println()
}

func operate(opIndex int, cIndex int, cVal uint8) {
	var buffer bytes.Buffer
	// fmt.Println("in Operate", opIndex, cIndex, cVal)
	op := '>'
	indexdiff := opIndex - cIndex
	if indexdiff < 0 {
		if indexdiff < -15 {
			op = '>'
			indexdiff = 30 - (indexdiff * -1)
		} else {
			op = '<'
			indexdiff = (indexdiff * -1)
		}
	} else {
		if indexdiff > 15 {
			op = '<'
			indexdiff = 30 - indexdiff
		} else {
			op = '>'
		}
	}
	//Print navigation
	// fmt.Println("in Operate: print navigation", string(op), indexdiff)
	for i := 0; i < indexdiff; i++ {
		buffer.WriteString(string(op))
		// fmt.Print(op)
	}

	// Print operation
	// fmt.Println("in Operate: print operation", stoneList[opIndex].value, cVal)
	diffValue := valDiff(cVal, stoneList[opIndex].value)
	// fmt.Println("in Operate: print operation(-:-, +:+)", diffValue)

	count := diffValue
	if count < 0 {
		count *= -1
	}
	for i := 0; i < count; i++ {
		if diffValue < 0 {
			// fmt.Print("-")
			buffer.WriteString("-")

		} else {
			// fmt.Print("+")
			buffer.WriteString("+")
		}
	}
	// fmt.Print(".")
	buffer.WriteString(".")

	stoneList[opIndex].value = cVal
	fmt.Print(buffer.String())
}

func initMap() {
	m = make(map[string]int)
	m[" "] = 0
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	m["D"] = 4
	m["E"] = 5
	m["F"] = 6
	m["G"] = 7
	m["H"] = 8
	m["I"] = 9
	m["J"] = 10
	m["K"] = 11
	m["L"] = 12
	m["M"] = 13
	m["N"] = 14
	m["O"] = 15
	m["P"] = 16
	m["Q"] = 17
	m["R"] = 18
	m["S"] = 19
	m["T"] = 20
	m["U"] = 21
	m["V"] = 22
	m["W"] = 23
	m["X"] = 24
	m["Y"] = 25
	m["Z"] = 26
}
