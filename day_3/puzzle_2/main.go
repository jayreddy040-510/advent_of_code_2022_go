package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "github.com/jayreddy040-510/advent_of_code_2022_go/day_3/utils"
)

var rucksackPrioMap = utils.MakePrioMaps()

func findCommonItem(group *[3]string) (rune, error) {
	// earlier used two sets but can instead use 1 set and flip the bool
	set := make(map[rune]bool)

	for _, ch := range group[0] {
		set[ch] = true
	}
	for _, ch := range group[1] {
		if set[ch] {
			set[ch] = false
		}
	}
	for _, ch := range group[2] {
		if val, exists := set[ch]; !val && exists {
			return ch, nil
		}
	}
	return -1, fmt.Errorf("Error finding common item among group of 3 elves: %v!", group)
}

func main() {
	file, err := os.Open("../puzzle_input.txt")
	if err != nil {
		log.Fatalf("Error opening puzzle input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numElves int
	var groupOfElves [3]string
	var totalPrioScore int

	for scanner.Scan() {
		line := scanner.Text()
		groupOfElves[numElves] = line
		numElves++
		if numElves == 3 {
			commonItem, err := findCommonItem(&groupOfElves)
			if err != nil {
				log.Fatalf("%v", err)
			}
			if points, exists := rucksackPrioMap[commonItem]; exists {
				totalPrioScore += points
			} else {
				log.Fatalf("Weird error, commonItem not found in rucksackPrioMap")
			}
			numElves = 0
		}
	}
	log.Printf("Total points: %d", totalPrioScore)
}
