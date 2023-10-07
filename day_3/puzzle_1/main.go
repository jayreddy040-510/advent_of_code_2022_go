package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "github.com/jayreddy040-510/advent_of_code_2022_go/day_3/utils"
)

var rucksackPrioMap = utils.MakePrioMaps()

func findCommonItemInRucksack(rucksack string) (rune, error) {
	midpoint := len(rucksack) / 2
	set := make(map[rune]bool)

	for idx, char := range rucksack {
		if idx <= midpoint-1 {
			set[char] = true
		} else {
			if set[char] {
				return char, nil
			}
		}
	}
	return -1, fmt.Errorf("Couldn't find matching item in second compartment of rucksack: %s", rucksack)
}

func main() {
	file, err := os.Open("../puzzle_input.txt")
	if err != nil {
		log.Fatalf("Error opening puzzle input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int

	for scanner.Scan() {
		line := scanner.Text()
		commonItem, err := findCommonItemInRucksack(line)
		if err != nil {
			log.Fatalf("%v", err)
		}
		count += rucksackPrioMap[commonItem]
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning puzzle input: %v", err)
	}
	log.Printf("Final count: %d", count)
}
