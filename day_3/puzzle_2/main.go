package main

import (
	"bufio"
	"fmt"
	"os"
	// rr "github.com/jayreddy040-510/advent_of_code_2022_go/day_3/utils"
)

func rucksackMapMaker(r string) map[rune]bool {

	retMap := make(map[rune]bool)

	for _, char := range r {
		retMap[char] = true
	}

	return retMap
}

func main() {
	// prioMap := rr.MakePrioMaps()
	retArray := make([]string, 100)

	f, err := os.Open("../puzzle_input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error returned: %v", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0

	var mapOne, mapTwo map[rune]bool = make(map[rune]bool), make(map[rune]bool)

	for scanner.Scan() {
		text := scanner.Text()
		if count == 0 {
			mapOne = rucksackMapMaker(text)
		} else if count == 1 {
			for _, char := range text {
				if mapOne != nil {
					if mapOne[char] {
						mapTwo[char] = true
					}
				} 
			}
		} else if count == 2 {
			for _, char := range text {
				if mapTwo != nil {
					if mapTwo[char] {
						retArray = append(retArray, string(char))
					}
				}
			}
			mapOne, mapTwo = make(map[rune]bool), make(map[rune]bool)
			count = 0
		}
		count++
	}

	returnSum := 0

	// for _, char := range retArray {
	// 	returnSum += prioMap[byte(char)]
	// }
	fmt.Println(retArray,returnSum)
}
