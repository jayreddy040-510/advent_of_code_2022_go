package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jayreddy040-510/advent_of_code_2022_go/day_3/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	prioMap := rucksackutils.MakePrioMaps()
	returnSum := 0
	data, err := os.ReadFile("puzzle_input.txt")
	check(err)

	dataString := string(data)

	rucksackSlice := strings.Split(dataString, "\n")

	for _, sack := range rucksackSlice {
		firstCompartmentContents := make(map[rune]bool)
		for i, char := range sack {
			if i < len(sack)/2 {
				firstCompartmentContents[char] = true
			} else {
				if firstCompartmentContents[char] {
					// fmt.Println(string(char), prioMap[char])
					returnSum += prioMap[char]
					break
				}
			}
		}
	}

	fmt.Println(returnSum)

}
