package main

import (
	"bufio"
	"fmt"
	"os"

	rr "github.com/jayreddy040-510/advent_of_code_2022_go/day_3/utils"
)

func rucksackMapMaker(r string) map[rune]bool {
	retMap := make(map[rune]bool)
	for _, char := range r {
		retMap[char] = true
	}
	return retMap
}

func main() {
	prioMap := rr.MakePrioMaps()
	retArray := make([]rune, 0)
	f, err := os.Open("../puzzle_input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error returned: %v", err))
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	count := 0
	var mapOne map[rune]bool = make(map[rune]bool)

	for scanner.Scan() {
		text := scanner.Text()
		if count == 0 {
			mapOne = rucksackMapMaker(text)
		} else if count == 1 {
			for k := range mapOne {
				if _, ok := rucksackMapMaker(text)[k]; !ok {
					delete(mapOne, k)
				}
			}
		} else if count == 2 {
			for k := range mapOne {
				if _, ok := rucksackMapMaker(text)[k]; ok {
					retArray = append(retArray, k)
				}
			}
			mapOne = make(map[rune]bool)
			count = -1
		}
		count++
	}

	returnSum := 0
	for _, char := range retArray {
		returnSum += prioMap[char]
	}
	fmt.Println(len(retArray), returnSum)
}
